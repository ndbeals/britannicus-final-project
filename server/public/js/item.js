inventoryID = 1;

$(document).ready(function () {

    inventoryID = parseInt($("#inventory_ID").val());
    console.log(inventoryID);


    $("#inventory_next").click(function () {
        inventoryID++;
        console.log(inventoryID);
        window.location.href = inventoryID;
    });

    $("#inventory_previous").click(function () {
        inventoryID--;
        window.location.href = inventoryID;
    });

    $("#inventory_ID").bind('change', function () {
        prodid = parseInt($("#inventory_ID").val());

        console.log("prodid")

        if (prodid > 0) {
            // ipcRenderer.send("set_inventory_page", pagenum);
            // changeinventory(prodid)
            inventoryID++;
            window.location.href = prodid;
        }
    });

    $("#inventory_delete").click( function(e) {
        e.preventDefault();
        console.log("delete");

        $.ajax({
            url: "/v1/inventory/" + inventoryID,
            dataType: 'json',
            type: 'DELETE',
            success: function (data) {
                // console.log("DATA POSTED SUCCESSFULLY" , data);
                alert(data.Message)
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message)
            }
        });
    })


    $("#inventory_updateform").submit( function(e) {
        e.preventDefault();
        // $("#inventory_ISBN").val(data.isbn)
        // $("#inventory_author").val(data.inventory_author)
        // $("#inventory_genre").val(data.inventory_genre)
        // $("#inventory_description").val(data.inventory_description)
        // $("#inventory_name").val(data.inventory_name)
        // $("#inventory_type").val(data.inventory_type)
        console.log("pub",$("#inventory_publisher").val());

        var data = {};
        data.inventory_id = inventoryID
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
            url: "/v1/inventory/" + inventoryID,
            dataType: 'json',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
            },
            error: function (jqXhr, textStatus, errorThrown) {
                console.log(errorThrown);
            }
        });
    });

    $("#inventory_condition").val(inventoryCondtion)
});