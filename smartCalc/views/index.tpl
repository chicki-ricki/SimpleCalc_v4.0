<!DOCTYPE html>

<html>
<head>
  <title>smartCalc</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="/static/img/Icon.png" type="image/x-icon" />
  <link rel="stylesheet" href="/static/css/index_style.css" />
</head>

<body>
  <header>
    <h1 class="logo">
      <a href="http://localhost:8080/calculate">Welcome to smartCalc</a>
    </h1>
    <div class="description">
      SmartCalc is a simple & powerful Go web calculator.
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

</body>
</html>
