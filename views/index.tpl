<!DOCTYPE html>

<html>
  	<head>
    	<title>Wash Sales</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
      <link rel="stylesheet" href="static/css/bootstrap.css">
      <link rel="stylesheet" href="static/css/master.css">
      <link rel="stylesheet" href="static/css/font-awesome.min.css">
	</head>

  	<body>
      <script type="text/x-handlebars" id="application">

            <nav class="navbar navbar-default" role="navigation">
            <div class="container">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">TaxForms</a>
            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">

            <ul class="nav navbar-nav navbar-right">
            {{#if isAuthenticated}}
              <li>{{#link-to "uploads.newUpload"}}{{controllers.auth.currentUser.email}} {{/link-to}}</li>
              <li><a href="#" {{action "logout"}}>Logout</a></li>
            {{else}}
              <li>{{#link-to "signIn"}}Sign In {{/link-to}}</li>
              <li>{{#link-to "signUp"}}Sign Up {{/link-to}}</li>
            {{/if}}
            </ul>
            </div><!-- /.navbar-collapse -->
            </div><!-- /.container-fluid -->
            </nav>
            <div class="container">
          {{outlet}}
        </div>
      </script>

      <script type="text/x-handlebars" id="signIn">
        <div class="panel panel-default">
        <div class="panel-heading"><h4 >Sign in</h4></div>
        <div class="panel-body">
        <form {{action "login"  on="submit"}} class="form-horizontal col-xs-12 col-sm-12 col-md-12  col-lg-12">

        <div class="form-group"><label class="control-label col-xs-3 col-sm-3 col-md-3  col-lg-3">Email</label>
        <div class="col-xs-7 col-sm-7 col-md-7  col-lg-7">
        {{input type="text" value=email type="text" id="email" class="form-control" required="true"}}
        </div>
        <div class="col-xs-1 col-sm-1 col-md-1  col-lg-1"></div></div>

        <div class="form-group"><label class="control-label col-xs-3 col-sm-3 col-md-3  col-lg-3">Password</label>
        <div class="col-xs-7 col-sm-7 col-md-7  col-lg-7">
        {{input type="text" value=password id="password" type="password" class="form-control" required="true" pattern=".{8,}" title="Passwords must be at least 8 characters."}}</div>

        <div class="col-xs-1 col-sm-1 col-md-1  col-lg-1"></div>
        </div>


        <div class="text-center col-xs-12 col-sm-12 col-md-12  col-lg-12">
        <button  type="submit" class="btn btn-success" value="Login">Login</button>
        </div>
        </form>
        </div>
        </div>

      </script>

      <script type="text/x-handlebars" id="signUp">

      <div class="panel panel-default">

      <div class=" panel-heading"><h4>Register an Account</h4></div>
      <div class="panel-body">

      <form {{action "register"  on="submit"}} id="register-form" class=" form-horizontal col-xs-12">

      <h4 class="text-center">Account Info</h4>
      <div class="form-group">

      <div class="form-group"><label class="control-label col-xs-3">Email*</label>
      <div class="col-xs-9">
      {{input value=email type="text" id="email" class="form-control" required="true" pattern=".+"}}
      </div>
      </div>

                      <div class="form-group"><label class="control-label col-xs-3">Password*</label>
                        <div class="col-xs-9">
                          {{input value=password type="password" id="password" class="form-control" required="true" pattern=".{8,}" title="Passwords must be at least 8 characters."}}
                        </div>
                      </div>

                      <div class="text-center col-xs-12"><button  type="submit" class="btn btn-success" value="Login">Register</button></div>

                      <div class="float-clear"></div>

                    </form>
                    <h4>*=Required</h4>
                  </div>
                </div>
                </div>

      </script>

      <script type="text/x-handlebars" id="uploads">
        <h2>Upload CSV File</h2>
        <table class="table">
        <thead>
          <th>File</th>
          <th>Download</th>
        </thead>
        <tbody>
          {{#each upload in model itemController="upload"}}
            <td>{{upload.file}}<td>
            <td>
              <div class="btn-group">
                <div class="btn btn-primary" {{action "download" upload.id}}>Download</div>
                <div class="btn btn-danger" {{action "delete" upload.id}}>Delete</div>
              </div>
            </td>
            </tbody>
          {{/each}}
        </table>

        <ul class="nav nav-tabs" role="tablist">
          <li role="presentation" class="active">{{#link-to "uploads.newUpload" }}New Upload{{/link-to}}</li>
          <li role="presentation">{{#link-to "uploads.transactions" selectedUpload.id }}Transactions{{/link-to}}</li>
        </ul>

        <!-- Tab panes -->
        <div class="tab-content">
        {{outlet}}
        </div>


        </script>


        <script type="text/x-handlebars" id="uploads/newUpload">
          <div role="tabpanel" class="tab-pane active">
            <h1>Upload Files ({{files.length}} in total)</h1>

            <ul class="counts">
            <li class="valid">Valid: {{validFiles.length}}</li>
            <li class="invalid">Invalid: {{invalidFiles.length}}</li>
            <li class="uploaded">Uploaded: {{uploadedFiles.length}}</li>
            <li class="deleted">Deleted: {{deletedFiles.length}}</li>
            </ul>

            <div class="buttons">
            <button class="btn" {{action "uploadAllFiles"}}>Upload All</button>
            <button class="btn" {{action "clearAllFiles"}}>Clear</button>
            <button class="btn" {{action "abortUpload"}}>Abort Upload</button>
            </div>

            {{#if uploadStatus.uploading}}
            <h3 class="uploading-percentage">Uploaded Percentage: {{uploadStatus.percentComplete}}%</h3>
            {{/if}}

            {{#view view.DragDrop}}

            {{#if uploadStatus.error}}
            <div class="error">An error occurred during the upload process. Please try again later.</div>
            {{/if}}

            {{#each controller.validFiles}}

            <div {{bind-attr class="className :file"}}>
            {{name}}
            <a class="remove" {{action "deleteFile" this}}>Discard.</a>
            {{view view.ImagePreview imageBinding="this"}}
            </div>

            {{/each}}

            {{view view.SingleInput}}

            {{/view}}

          </div>
          </script>


        <script type="text/x-handlebars" id="uploads/transactions">
        <div role="tabpanel" class="tab-pane" >

          <br>
          <h2>Transactions for {{currentUpload.file}}</h2>
          <table class="table">
          <thead>
          <th>Date</th>
          <th>Ticker</th>
          <th>Action</th>
          <th>Quantity</th>
          <th>Price</th>
          </thead>
          <tbody>
          {{#if currentUpload}}
          {{#each transaction in currentUpload.transactions}}
          <td>{{transaction.date}}</td>
          <td>{{transaction.ticker}}</td>
          <td>{{transaction.action}}</td>
          <td>{{transaction.quantity}}</td>
          <td>{{transaction.price}}</td>
          </td>
          </tbody>
          {{/each}}
          {{/if}}
          </table>

        </div>
        </script>










      <script src="static/js/jquery-1.10.2.js"></script>
      <script src="static/js/bootstrap.js"></script>
      <script src="static/js/handlebars-v1.3.0.js"></script>
      <script src="static/js/ember-1.8.1.js"></script>
      <script src="static/js/ember-data.js"></script>
      <script src="static/js/ember-droplet-mixin.js"></script>
      <script src="static/js/ember-droplet-view.js"></script>
      <script src="static/js/app.js"></script>
      <script src="static/js/authController.js"></script>
	  </body>
</html>
