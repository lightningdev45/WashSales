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
        <div class="container">
          {{outlet}}
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
	  </body>
</html>
