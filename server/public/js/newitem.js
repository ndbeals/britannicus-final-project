productID = 1;

$(document).ready(function () {

    productID = parseInt($("#inventory_ID").val());
    console.log(productID);


    $("#product_next").click(function () {
        productID++;
        populateProduct(productID)
        // console.log(productID);
        // window.location.href = productID;
    });

    $("#product_previous").click(function () {
        productID--;
        populateProduct(productID)
        // window.location.href = productID;
    });

    $("#product_ID").bind('change', function () {
        prodid = parseInt($("#product_ID").val());

        console.log("prodid",prodid)

        if (prodid > 0) {
            // ipcRenderer.send("set_inventory_page", pagenum);
            // changeinventory(prodid)
            // productID++;
            populateProduct(prodid)
            // window.location.href = prodid;
        }
    });


    $("#inventory_updateform").submit( function(e) {
        e.preventDefault();
        // $("#inventory_ISBN").val(data.isbn)
        // $("#inventory_author").val(data.inventory_author)
        // $("#inventory_genre").val(data.inventory_genre)
        // $("#inventory_description").val(data.inventory_description)
        // $("#inventory_name").val(data.inventory_name)
        // $("#inventory_type").val(data.inventory_type)
        // console.log("pub",$("#inventory_publisher").val());

        var data = {};
        data.inventory_id = productID
        data.product_id = productID
        data.inventory_condition = parseInt($("#inventory_condition").val())
        data.inventory_amount = parseInt($("#inventory_amount").val())
        data.inventory_price = parseFloat($("#inventory_price").val())
        data.inventory_note = $("#inventory_note").val()
        // data.inventory_description = $("#inventory_description").val()
        // data.inventory_name = $("#inventory_name").val()
        // data.inventory_type = $("#inventory_type").val()
        
        console.log($("#inventory_condition").val())

        test = $("#inventory_updateform")

        console.log(JSON.stringify(data));


        $.ajax({
            url: "/v1/inventory",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
                window.location.href="/inventory/get/"+data.id
            },
            error: function (jqXhr, textStatus, errorThrown) {
                console.log(errorThrown);
            }
        });
    });

    $("#inventory_condition").val(inventoryCondtion)

    populateProduct(1)
});

function populateProduct( id ) {
    $.get("/v1/product/"+id, function(data) {
        console.log(data)
        if (data !== null) {
            productID = id
            $("#product_ID").val(id)

            $("#inventory_isbn").val(data.isbn)
            $("#inventory_author").val(data.product_author)
            $("#inventory_genre").val(data.product_genre)
            $("#inventory_publisher").val(data.product_publisher)
            $("#inventory_description").val(data.product_description)
            $("#inventory_name").val(data.product_name)
            $("#inventory_type").val(data.product_type)
        }
    })
}