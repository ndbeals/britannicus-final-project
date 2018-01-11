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
        // $("#order_ISBN").val(data.isbn)
        // $("#order_author").val(data.order_author)
        // $("#order_genre").val(data.order_genre)
        // $("#order_description").val(data.order_description)
        // $("#order_name").val(data.order_name)
        // $("#order_type").val(data.order_type)
        console.log("pub",$("#order_publisher").val());

        var data = {};
        data.order_id = orderID
        // data.order_condition = parseInt($("#order_condition").val())
        // data.order_amount = parseInt($("#order_amount").val())
        // data.order_price = parseFloat($("#order_price").val())
        // data.order_note = $("#order_note").val()
        // data.order_description = $("#order_description").val()
        // data.order_name = $("#order_name").val()
        // data.order_type = $("#order_type").val()
        
        console.log($("#order_condition").val())

        test = $("#order_updateform")

        console.log(JSON.stringify(data));


        $.ajax({
            url: "/v1/order/" + orderID,
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

    // $("#order_condition").val(orderCondtion)
    updateList();
    console.log(orderID);
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
    // console.log("tast",item)
    // var row = '<tr><td colspan="3">' + '</td><td><div id="#collapse' + order_id +' class="collapse in"' +
    //     item.inventory_id + "</div></td><td>" +
    //     item.product.product_name + "</td><td>" +
    //     "" + "</td><td>" +

    //     +'</td></div></tr>'

    console.log(item)
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

    row = $(row).appendTo(table)//table.append(row);
    // row.toggle(false)


    // $('#btndel' + order_id).click(function () {
    //     var parent = $(this).parent().parent();
    //     $.ajax({
    //         url: "/v1/order/" + order_id,
    //         dataType: 'json',
    //         type: 'DELETE',
    //         success: function (data) {
    //             if (data !== null) {
    //                 parent.remove()
    //             }
    //         },
    //         error: function (data, textStatus, errorThrown) {
    //             // alert(data.responseJSON.Message)
    //             console.log(data);
    //         },
    //     });
    // });

    return row;
}