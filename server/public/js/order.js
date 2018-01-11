orderID = 1;

$(document).ready(function () {

    orderID = parseInt($("#order_ID").val());
    console.log(orderID);


    $("#order_next").click(function () {
        orderID++;
        console.log(orderID);
        window.location.href = orderID;
    });

    $("#order_previous").click(function () {
        orderID--;
        window.location.href = orderID;
    });

    $("#order_ID").bind('change', function () {
        prodid = parseInt($("#order_ID").val());

        console.log("prodid")

        if (prodid > 0) {
            // ipcRenderer.send("set_order_page", pagenum);
            // changeorder(prodid)
            orderID++;
            window.location.href = prodid;
        }
    });

    $("#order_delete").click( function(e) {
        e.preventDefault();
        console.log("delete");

        $.ajax({
            url: "/v1/order/" + orderID,
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


    $("#order_updateform").submit( function(e) {
        e.preventDefault();

        var data = {};
        data.order_id = orderID

        $.ajax({
            url: "/v1/order/" + orderID,
            dataType: 'json',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });

    // $("#order_condition").val(orderCondtion)
    updateList();
});

function updateList() {
    $.get( "/v1/order/"+orderID , function( data ) {
        if (data !== null ){
            table = $("#orderlisttable")

            for (var i=0; i<data.item_list.length; i++){
                addOrderItem( table , orderID, data.item_list[i] )
            }
        }
    })
}

function addOrderItem(table, order_id, item) {
    var row = `
    <tr>
        <td>
            <div>
                `+ item.inventory_id + `
            </div>
        </td>
        <td>
            <div>
               `+ item.product.product_name + `
            </div>
        </td>
        <td>
            <div>
                `+ item.product.product_author + `
            </div>
        </td>
        <td>
            <div>
                `+ item.order_amount + `
            </div>
        </td>
        <td>
            <div>
                `+ item.item_price + `
            </div>
        </td>
        <td>
            <div>
                `+ item.inventory_condition + `
            </div>
        </td>
        <td colspan='3'></td>
    </tr>`

    row = $(row).appendTo(table)

    return row;
}