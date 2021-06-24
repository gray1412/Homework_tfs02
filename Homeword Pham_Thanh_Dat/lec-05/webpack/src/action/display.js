function display(data) {
  $("td").remove();
  for (var i = 0; i < data.length; i++) {
    $("#myTable").append(
      "<tr><td>" +
        data[i].Id +
        "</td>" +
        data[i].Id +
        "<td>" +
        data[i].name +
        "</td><td>" +
        data[i].address +
        "</td><td>" +
        data[i].phone +
        "</td><td>" +
        data[i].age +
        "</td></tr>"
    );
  }
}
export { display };