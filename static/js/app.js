window.App = Ember.Application.create();
App.ApplicationAdapter = DS.RESTAdapter.extend({

});

App.ApplicationSerializer = DS.RESTSerializer.extend({

});

App.Router.map(function() {

    this.resource("uploads",{path:"/"},function(){
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

  App.UploadsTransactionsController=Ember.ArrayController.extend({

  })

  App.UploadsTransactionsRoute=Ember.Route.extend({
    model:function(params){
      console.log(params.uploadId)
      return this.store.getById("upload",params.uploadId).get("transactions")
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
