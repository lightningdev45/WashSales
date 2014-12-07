var App = Ember.Application.create();
App.ApplicationAdapter = DS.RESTAdapter.extend({

});

App.ApplicationSerializer = DS.RESTSerializer.extend({

});

App.Router.map(function() {

    this.route("room",{path:"/room/:roomId"})



  });

  App.Room=DS.Model.extend({
    roomId:DS.attr()
  })
  App.RoomController=Ember.Controller.extend({

  })

  App.RoomRoute=Ember.Route.extend({
    model:function(params){
      this.set("roomId",params.roomId)
      return {}
    },
    setupController:function(controller,model){
      this._super(controller, model);
      controller.set("roomId",this.get("roomId"))
    }
  })

  App.RoomView=Ember.View.extend({
    didInsertElement:function(){
      console.log(this.get("controller"))
      console.log(this.get("controller.roomId"))
      this.set("roomSocket", new WebSocket("ws://10.0.0.2:8080/room/"+this.get("controller.roomId")))
      var roomSocket=this.get("roomSocket")
      roomSocket.onmessage = function (event) {
        console.log("socket")
        console.log(event)
        $("#messages").append("<p>"+event.data+"</p>")
      }
    },
    willDestroyElement:function(){
      //this.get("roomSocket").close()
    }

  })

  App.ChatTextFieldView=Ember.TextField.extend({
    input:function(){
      this.get("parentView").get("roomSocket").send(this.get("value"))

    }
  })
