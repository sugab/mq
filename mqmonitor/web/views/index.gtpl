<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8" />

		<title>MQ Monitor</title>

		<!-- res/jquery -->
		<script src="/res/jquery/jquery-2.1.3.min.js"></script>

		<!-- res/bootstrap -->
		<link rel="stylesheet" href="/res/bootstrap/css/bootstrap.min.css">
		<link rel="stylesheet" href="/res/bootstrap/css/bootstrap-theme.min.css">
		<script src="/res/bootstrap/js/bootstrap.min.js"></script>

		<!-- res/kendoui -->
		<link rel="stylesheet" href="/res/kendoui/styles/kendo.common.min.css" />
		<link rel="stylesheet" href="/res/kendoui/styles/kendo.flat.min.css" />
		<link rel="stylesheet" href="/res/kendoui/styles/kendo.dataviz.min.css" />
		<link rel="stylesheet" href="/res/kendoui/styles/kendo.dataviz.flat.min.css" />
		<script src="/res/kendoui/js/kendo.all.min.js"></script>

		<!-- css/font-awesome -->
		<link rel="stylesheet" href="/res/font-awesome/css/font-awesome.min.css">

		<!-- res/main -->
		<link rel="stylesheet" href="/res/main/main.css">
		<script src="/res/main/main.js"></script>
	</head>

	<body>
		<div class="container-fluid main no-padding">
			<div class="page-header">
				<div class="pull-left">
					<img class="logo" src="/res/images/logo.png" />
				</div>
				<div class="pull-left">
					<h1>MQ Monitor</h1>
				</div>
				<div class="clearfix"></div>
			</div>

			<div class="col-md-6 section section-nodes">
				<div class="panel panel-primary">
					<div class="panel-heading">
						<i class="fa fa-cogs"></i> Nodes Information
					</div>
					<div class="panel-body">
						<div class="grid"></div>
					</div>
				</div>
			</div>

			<div class="loader">
				<img src="/res/images/495.gif" />
			</div>
		</div>
	</body>
</html>