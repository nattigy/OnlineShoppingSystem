// delete
var index,
  table = document.getElementById("table");
var index1,
  table1 = document.getElementById("table1");

// previous post
for (var i = 1; i < table.rows.length; i++) {
  table.rows[i].cells[4].onclick = function() {
    var c = confirm("Do you want to Remove?");
    if (c === true) {
      index = this.parentElement.rowIndex;
      table.deleteRow(index);
    }
  };
}
// new post
for (var i = 1; i < table1.rows.length; i++) {
  table1.rows[i].cells[2].onclick = function() {
    var alert = confirm("Do you want to Remove?");
    if (alert === true) {
      index1 = this.parentElement.rowIndex;
      table1.deleteRow(index1);
    }
  };
}
