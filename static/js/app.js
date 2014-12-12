window.App = Ember.Application.create();
App.ApplicationAdapter = DS.RESTAdapter.extend({

});

App.ApplicationSerializer = DS.RESTSerializer.extend({

});

App.Router.map(function() {
    this.route("signIn",{path:"/"})
    this.route("signUp")
    this.resource("uploads",function(){
      this.route("newUpload")
      this.route("transactions",{path:"/transactions/:uploadId"})

    })



  });

  App.Upload=DS.Model.extend({
    file:DS.attr(),
    transactions:DS.hasMany("transaction",{async:true})
  })

  App.Transaction=DS.Model.extend({
    ticker:DS.attr(),
    action:DS.attr(),
    quantity:DS.attr(),
    price:DS.attr(),
    date:DS.attr(),
    upload:DS.belongsTo("upload",{async:true})
  })

  App.User=DS.Model.extend({
    email:DS.attr(),
    password:DS.attr()
  })

  App.UploadsTransactionsController=Ember.ArrayController.extend({

  })

  App.UploadsTransactionsRoute=Ember.Route.extend({
    model:function(params){
      console.log(params.uploadId)
      return this.store.getById("upload",params.uploadId).get("transactions")
    }

  })



  App.SignInRoute = Ember.Route.extend({
    activate:function(){
      console.log("hi")
      var controller=this;
      if(this.controllerFor("auth").get("isAuthenticated")){
        this.controllerFor("alert").send("showAlert","You are already signed in!","alert alert-warning alert-dismissible","devise-alert")

        this.controllerFor("auth").get("currentUser").get("account").then(function(account){
          controller.transitionTo("upload.newUpload")
        })


      }
    },
    setupController: function(controller, model){
      controller.setProperties({
        password: "",
        errorMsg: ""
      });

    },

    actions:{
      login: function(){
        //console.log(this.get("controller"))
        this.controllerFor("auth").login(this)
      },
      cancel:function(){
        this.transitionTo('entities.index')
      }
    }
  })

  App.ApplicationController=Ember.Controller.extend({
    needs:["auth"],
    isAuthenticated: Em.computed.alias("controllers.auth.isAuthenticated"),
    actions:{
      logout:function(){
        this.get("controllers.auth").logout()
      }
    }
  })

  App.SignUpRoute = Ember.Route.extend({

    actions:{
      register:function(){
        var route= this
            $.ajax({
              url: "/users_create",
              type: "POST",
              data:{
                "email": route.controller.get("email"),
                "password": route.controller.get("password")
              },
              success: function(data){
                var auth=route.controllerFor("auth")
                if(data.user){
                  var serializedCurrentUser=auth.store.serializerFor("User").normalize(App.User,data.user)
                  auth.set("currentUser",auth.store.push("User",serializedCurrentUser))
                }
                $('meta[name="csrf-token"]').attr('content', data['csrf-token'])
                $('meta[name="csrf-param"]').attr('content', data['csrf-param'])
                route.controllerFor("signUp").setProperties({password:"",email:""})
                //route.controllerFor("alert").send("showAlert","You have successfully registered your account for our site!","alert alert-success alert-dismissible","devise-alert")
                route.transitionTo('uploads.newUpload')
                },
                error: function(jqXHR, textStatus, errorThrown){
                  if(jqXHR.responseJSON.errors){
                    if(jqXHR.responseJSON.errors.email){
                      //route.controllerFor("alert").send("showAlert","That email "+jqXHR.responseJSON.errors.email+".","alert alert-danger alert-dismissible","devise-alert")
                    }
                    else if(jqXHR.responseJSON.errors.profile_name){
                      //route.controllerFor("alert").send("showAlert","That profile name "+jqXHR.responseJSON.errors.profile_name+".","alert alert-danger alert-dismissible","devise-alert")
                    }
                    else{
                      //route.controllerFor("alert").send("showAlert","There was an error.  Please register again or contact support.","alert alert-danger alert-dismissible","devise-alert")
                    }
                  }
                  else{
                    //route.controllerFor("alert").send("showAlert","There was an error.  Please register again or contact support.","alert alert-danger alert-dismissible","devise-alert")
                  }
                }
              })

            }
          }



    })

  App.UploadsController=Ember.ArrayController.extend({
    selectedUpload:function(){
      console.log(this.get("model.firstObject.id"))
      return this.get("model.firstObject")
    }.property("model.[]"),
    actions:{
      download:function(id){
        location.href="/download_csv/"+id
      },
      delete:function(id){
        var controller=this;
        jQuery.ajax({
          url:"/delete_csv/"+id,
          method:"POST",
          success:function(){
            alert("success")
            controller.store.unloadRecord(controller.store.getById("upload",id))
          },
          error:function(){
            alert("error")
          }
        })
      },
      getTransactions:function(id){
        var controller=this;
        var currentUpload=controller.store.getById("upload",id)
        controller.set("currentUpload",currentUpload)
        jQuery.ajax({
          url:"/get_transactions/"+id,
          method:"GET",
          success:function(data){
            alert("success")
            console.log(data.transactions)
            controller.store.pushMany("transaction",data.transactions)
            console.log(controller.store.all("transaction"))
            console.log(currentUpload)
          },
          error:function(){
            alert("error")
          }
        })
      }
    }

  })


  App.UploadsNewUploadView=Ember.View.extend({

    DragDrop: DropletView.extend()
  })

  App.UploadsNewUploadController=Ember.ArrayController.extend(DropletController,{
    dropletUrl: 'upload',
    dropletOptions: {
      fileSizeHeader: true,
      useArray: false
    },
    /**
    * Specifies the valid MIME types. Can used in an additive fashion by using the
    * property below.
    *
    * @property mimeTypes
    * @type {Array}
    */
    mimeTypes: ['image/bmp','text/csv'],
    /**
    * Apply this property if you want your MIME types above to be appended to the white-list
    * as opposed to replacing the white-list entirely.
    *
    * @property concatenatedProperties
    * @type {Array}
    */
    concatenatedProperties: ['mimeTypes'],
    /**
    * @method didUploadFiles
    * @param response {Object}
    * @return {void}
    */
    didUploadFiles: function(response) {
      this.store.push("upload",response)
    }
  })

  App.UploadsRoute=Ember.Route.extend({
    model:function(params){
      return this.store.find("upload")
    },
    afterModel:function(){
      this.transitionTo("uploads.newUpload")
    }
  })

  App.UploadController=Ember.ObjectController.extend({
    needs:["uploads"],
    isSelected:Ember.computed.equal("id","controllers.uploads.currentUpload.id"),
  })

  App.UploadView=Ember.View.extend({
    didInsertElement:function(){
      console.log("hi")
    },

    click:function(){
      this.get("parentView.controller").set("selectedUpload",this.get("controller"))
    },
    willDestroyElement:function(){
      //this.get("roomSocket").close()
    }

  })
