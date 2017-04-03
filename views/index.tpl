<!DOCTYPE html>

<html>
<head>
  <title>中投天诚数据报告</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="/static/css/bootstrap.css" rel='stylesheet' type='text/css'/>
  <style type="text/css" media="screen">
    body{
      background-image: url('/static/img/bg1.jpg');
      font-weight: bold;
    }
    iframe{
      color: #fff;
    }
  </style>
  <script src="/static/js/jquery.js"></script>
  <script src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/Chart.min.js"></script>
  <script src="/static/js/home.js"></script>
  <script type="text/javascript">
  </script>
</head>

<body>
   <h1 class="logo" style="text-align: center">中投天诚报表</h1>
   <div class="col-md-4 col-md-offset-8">
        <a href="/admin" class="btn btn-info">进入后台</a> 
        <div class="btn-group">
         <button type="button" class="btn btn-info dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded>
           双流店<span class="caret"></span>  
         </button>
         <ul class="dropdown-menu">
           <li><a href="#" title="">双流店</a></li>
           <li><a href="#" title="">青白江店</a></li>
         </ul>
    </div>
  </div>
  <header class="container">
      <div class="col-md-6">
        <h3 style="text-align: center">每周数据</h3>
        <canvas id="chart1" width="400" height="260"></canvas>
      </div>
       <div class="col-md-6">
       <h3 style="text-align: center;">每日报表</h3>
       <canvas id="chart2" width="400" height="260"></canvas>
      </div>
  </header>
  <footer style="text-align: center;">
    <h2 style="color: #fff">{{date .Time "Y年m月d日  H时:i分"}}</h2>
  </footer>
  <div class="backdrop"></div>
</body>
</html>
