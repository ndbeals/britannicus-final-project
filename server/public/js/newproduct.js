productID = 1;

$(document).ready(function () {

    productID = parseInt($("#product_ID").val());
    console.log(productID);


    $("#product_next").click(function () {
        console.log("next");
        changeProduct(productID + 1)
            // changeProduct(productID - 1);
        
    });

    $("#product_previous").click(function () {
        changeProduct(productID - 1)
    });

    $("#product_ID").bind('keyup input', function () {
        prodid= parseInt($("#product_ID").val());

        if (prodid > 0) {
            // ipcRenderer.send("set_inventory_page", pagenum);
            changeProduct(prodid)
        }
    });


    $("#product_updateform").submit( function(e) {
        e.preventDefault();
        var data = {};
        data.product_id = -1
        data.product_isbn = $("#product_isbn").val()
        data.product_publisher = $("#product_publisher").val()
        data.product_author = $("#product_author").val()
        data.product_genre = $("#product_genre").val()
        data.product_description = $("#product_description").val()
        data.product_name = $("#product_name").val()
        data.product_type = $("#product_type").val()
        


        $.ajax({
            url: "/v1/product",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                // console.log("DATA POSTED SUCCESSFULLY", data);
                window.location.href="/product/get/"+data.id
            },
            error: function (jqXhr, textStatus, errorThrown) {
                console.log(errorThrown);
            }
        });
    });

    
});

function changeProduct(id) {
    productID = id;
    $("#product_ID").val(id)

    if ( id > 0) {

        $.get("/v1/product/" + productID, function (data) {
            console.log("still",data);
            if (data !== null) {
                data = data.data
                $("#product_ISBN").val(data.isbn)
                $("#product_author").val(data.product_author)
                $("#product_genre").val(data.product_genre)
                $("#product_description").val(data.product_description)
                $("#product_name").val(data.product_name)
                $("#product_type").val(data.product_type)
                $("#product_publisher").val(data.product_publisher)
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
    }
}
