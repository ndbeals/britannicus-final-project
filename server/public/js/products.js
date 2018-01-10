$(document).ready(function () {
    $("#productsFilterInput").on("keyup", function () {
        console.log("GA");
        var value = $(this).val().toLowerCase();
        $("#productsFilterTable tr").filter(function () {
            $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
        });
    });


    $.get("http://localhost:9000/v1/products/" + 1 + "/20", function (data) {
        if (data !== null) {

            var table = $("#productsFilterTable");

            table.empty();

            for (var i = 0; i < data.length; i++) {
                item = data[i];

                addProduct(table, item.product_id, item.isbn, item.product_name, item.product_author, item.product_genre, item.product_publisher, item.product_type, item.product_description)
                // AddItem(item.inventory_id, item.product.product_name, item.product.product_genre, item.amount, item.product.product_price, item.inventory_condition);
            }
        }
    })
    .done(function () {
        //alert("second success");
    })
    .fail(function () {
        //alert("error");
    })
    .always(function (data) {
        //alert("finished" + data);
    });
});

function addProduct(table, productID, ISBN, productName, author, genre, publisher, productType, description) {

    var row = "<tr><td>" + 
        productID + "</td><td>" +
        ISBN + "</td><td>" +
        productName + "</td><td>" +
        author + "</td><td>" +
        genre + "</td><td>" +
        publisher + "</td><td>" +
        productType + "</td><td>" +
        description + "</td></tr>"
        // </tr>

    table.append(row);
}