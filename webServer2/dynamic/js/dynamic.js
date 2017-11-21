
$(document).ready(function() {
      var dateNtime = new Date();
      var year = dateNtime.getFullYear();
      var month = dateNtime.getMonth() + 1;
      var day = dateNtime.getDate();
      var hours = dateNtime.getHours();
      var minute = dateNtime.getMinutes();
      var second = dateNtime.getSeconds();

      var timeStr = hours + ":" + minute + ":" + second;
      var dateStr = month + "/" + day + ", " + year + " " + timeStr;

      $('.demo').html("Today is " + dateStr);
})
