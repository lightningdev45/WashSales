window.App = Ember.Application.create();
App.ApplicationAdapter = DS.RESTAdapter.extend({

});

App.ApplicationSerializer = DS.RESTSerializer.extend({

});

App.Router.map(function() {

    this.route("upload",{path:"/"})



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

  App.UploadController=Ember.ArrayController.extend(DropletController,{
    currentUpload:null,
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
    },

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
      console.log(response);
      this.store.push("upload",response)
    }

  })

  App.UploadRoute=Ember.Route.extend({
    model:function(params){
      return this.store.find("upload")
    }
  })

  App.UploadView=Ember.View.extend({
    didInsertElement:function(){
      console.log("hi")
    },
    willDestroyElement:function(){
      //this.get("roomSocket").close()
    },

    DragDrop: DropletView.extend()

  })
