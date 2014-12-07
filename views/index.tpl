<!DOCTYPE html>

<html>
  	<head>
    	<title>Chatroom</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
      <link rel="stylesheet" href="static/css/bootstrap.css">
      <link rel="stylesheet" href="static/css/master.css">
      <link rel="stylesheet" href="static/css/font-awesome.min.css">
	</head>

  	<body>
      <script type="text/x-handlebars" id="application">
        <div class="container">
          {{outlet}}
        </div>
      </script>

      <script type="text/x-handlebars" id="index">
        {{#link-to "room" 1}}Room 1{{/link-to}}
        {{#link-to "room" 2}}Room 2{{/link-to}}
        {{#link-to "room" 3}}Room 3{{/link-to}}
        {{#link-to "room" 4}}Room 4{{/link-to}}
        {{#link-to "room" 5}}Room 5{{/link-to}}
        {{#link-to "room" 6}}Room 6{{/link-to}}
        {{#link-to "room" 7}}Room 7{{/link-to}}
        {{#link-to "room" 8}}Room 8{{/link-to}}
        {{#link-to "room" 9}}Room 9{{/link-to}}
        {{#link-to "room" 10}}Room 10{{/link-to}}
      </script>

      <script type="text/x-handlebars" id="room">
        <h1>Room {{roomId}}</h1>
        <br>
        <br>
        {{view "chat-text-field"}}
        <br>
        Messages
        <div id="messages">
        </div>
      </script>

      <script src="static/js/jquery-1.10.2.js"></script>
      <script src="static/js/bootstrap.js"></script>
      <script src="static/js/handlebars-v1.3.0.js"></script>
      <script src="static/js/ember-1.8.1.js"></script>
      <script src="static/js/ember-data.js"></script>
      <script src="static/js/app.js"></script>
	  </body>
</html>
