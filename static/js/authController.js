App.AuthController = Ember.ObjectController.extend({
  currentUser:null,
  isAuthenticated: function() {
    var me=this;
    if(this.get("currentUser")){
      return true
    }
    else{
      $.getJSON("/current_user").then( function( data ) {
        if(data){
          var serializedCurrentUser=me.store.serializerFor("User").normalize(App.User,data.user)
          me.set("currentUser",me.store.update("User",serializedCurrentUser))
        }
      })
    }
  }.property('currentUser'),
  login: function(route){
    var me = this
    $.ajax({
      url: "/sign_in",
      type: "POST",
      data:{"email": route.controller.get("email"),
      "password": route.controller.get("password")},
      success:function(data){
        console.log("Login Msg #{data.user.dummy_msg}")
        $('meta[name="csrf-token"]').attr('content', data['csrf-token'])
        $('meta[name="csrf-param"]').attr('content', data['csrf-param'])
        if(data.user){
          var serializedCurrentUser=me.store.serializerFor("User").normalize(App.User,data.user)
          me.set("currentUser",me.store.update("User",serializedCurrentUser))
        }
        route.controllerFor("signIn").setProperties({password:"",email:""})

          route.transitionTo("uploads.newUpload")


      },
      error:function (jqXHR, textStatus, errorThrown){
        if(jqXHR.responseJSON.error)
          {
            //route.controllerFor("alert").send("showAlert",jqXHR.responseJSON.error,"alert alert-danger alert-dismissible","devise-alert")
            }
          if(jqXHR.status==401)
            {route.controllerFor('login').set("errorMsg", "That email/password combo didn't work.  Please try again")}
            else if(jqXHR.status==406)
              { route.controllerFor('login').set("errorMsg", "Request not acceptable (406):  make sure Devise responds to JSON.")}
              else
                {console.log("Login Error: #{jqXHR.status} | #{errorThrown}")}
              }

            })

          },
          logout:function(){
            console.log("Logging out...")
            me = this
            //token = $('meta[name="csrf-token"]').attr('content')
            //$.ajaxPrefilter(function(options, originalOptions, xhr){
            //xhr.setRequestHeader('X-CSRF-Token', token)
            //})
            $.ajax({
              url: "/sign_out",
              type: "DELETE",
              dataType: "json",
              async:true,
              success:function (data, textStatus, jqXHR){
                $('meta[name="csrf-token"]').attr('content', data['csrf-token'])
                $('meta[name="csrf-param"]').attr('content', data['csrf-param'])
                console.log("Logged out on server")
                me.set('currentUser', null)
                //me.get("controllers.alert").send("showAlert","You have successfully logged out!","alert alert-success alert-dismissible","devise-alert")
                me.transitionToRoute("uploads.newUpload")
                },
                error:function (jqXHR, textStatus, errorThrown){
                  //route.controllerFor("alert").send("showAlert","There was an error.  Please try logging out again and/or contact support.","alert alert-danger alert-dismissible","devise-alert")
                }
              })
            }
      })
