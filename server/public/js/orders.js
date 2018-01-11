orderPage = 1;

$(document).ready(function () {
    orderPage = parseInt($("#order_page").val());


    $("#order_next").click(function () {
        if (!changePage(orderPage + 1)) {
            changePage(orderPage - 1);
        }
    });

    $("#order_previous").click(function () {
        if (!changePage(orderPage - 1)) {
            changePage(orderPage + 1);
        }
    });

    $("#order_page").bind('keyup input', function () {
        pagenum = parseInt($("#order_page").val());

        if (pagenum > 0) {
            oldpage = orderPage;

            if (!changePage(pagenum)) {
                changePage(oldpage);
            }
        }
    });


    $("#ordersFilterInput").on("keyup", function () {
        var value = $(this).val().toLowerCase();

        if (value == "") {
            changePage(orderPage);
        } else {
            results = 0
            $("#ordersFilterTable tr").filter(function () {
                if ($(this).is(":visible")) {
                    $(this).toggle(false)
                }
                if (results < 25) {
                    show = $(this).text().toLowerCase().indexOf(value) > -1
                    $(this).toggle(show)
                    if (show) {
                        results++;
                    }
                }
            });
        }
    });

    populateOrder(1);
});

function addOrder(table, order_id, customer_id , customer_name , total_price) {

    var row = '<tr class="accordion-toggle" data-toggle="collapse" data-target=".collapse' + order_id + '"><td>' +
        order_id + "</td><td>" +
        customer_id + "</td><td>" +
        customer_name + "</td><td colspan='6'>" +
        "" + '</td><td>' +
        total_price.toFixed(2) + "</td><td>" +
        '<a href="/order/get/' + order_id + '"><button type="button" class="btn btn-primary btn-block tbl-btn">Edit</button></a></td><td><button id="btndel' + order_id + '" type="button" class="btn btn-danger btn-block tbl-btn">Delete</button></td></tr>'

    row = $(row).appendTo(table)//table.append(row);



    $('#btndel' + order_id).click(function () {
        var parent = $(this).parent().parent();
        $.ajax({
            url: "/v1/order/" + order_id,
            dataType: 'json',
            type: 'DELETE',
            success: function (data) {
                if (data !== null) {
                    parent.remove()
                }
            },
            error: function (data, textStatus, errorThrown) {
                // alert(data.responseJSON.Message)
                console.log(data);
            },
        });
    });

    return row;
}

function addOrderItem(table, order_id, item) {
    // console.log("tast",item)
    // var row = '<tr><td colspan="3">' + '</td><td><div id="#collapse' + order_id +' class="collapse in"' +
    //     item.inventory_id + "</div></td><td>" +
    //     item.product.product_name + "</td><td>" +
    //     "" + "</td><td>" +

    //     +'</td></div></tr>'
    var row = `<tr>
        <div id="collapse1" class="collapse in">
                <td colspan="3">
                        Details 1 Details 2 Details 3
                        </td>
                        <td>test</td>
                        </div>
            </tr>`
            
    var row = `
    <tr>
        <td colspan=3></td>
        <td>
            <div class="collapse`+order_id+` collapse">
                `+item.inventory_id+`
            </div>
        </td>
        <td>
            <div class="collapse`+ order_id +` collapse">
               `+ item.product.product_name +`
            </div>
        </td>
        <td>
            <div class="collapse`+ order_id +` collapse">
                `+ item.inventory_condition +`
            </div>
        </td>
        <td>
            <div class="collapse`+ order_id +` collapse">
                `+ item.order_amount +`
            </div>
        </td>
        <td>
            <div class="collapse`+ order_id +` collapse">
                `+ item.item_price +`
            </div>
        </td>
        <td>
            <div class="collapse`+ order_id +` collapse">
                `+ item.note +`
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

function populateOrder(page, hide) {
    $.get("/v1/orders/" + page + "/15", function (data) {
        if (data !== null) {
            // console.log(data);
            var table = $("#ordersFilterTable");
            
            // table.empty();
            
            for (var i = 0; i < data.length; i++) {
                item = data[i];
                console.log(item)

                row = addOrder(table, item.order_id, item.customer.customer_id, item.customer.first_name + " " + item.customer.last_name, item.total_price)

                if (hide == true) {
                    row.toggle(false)
                    // changePage(orderPage);
                }

                if (item.item_list.length > 0) {
                    for (var ii = 0; ii < item.item_list.length; ii++) {
                        // console.log("wat",ii,table, item.item_list[ii])
                        row = addOrderItem(table, item.order_id, item.item_list[ii])

                        if (hide == true) {
                            row.toggle(false)
                            // changePage(orderPage);
                        }
                    }
                }
            }

            populateOrder(page + 1, true);
        }
        else {
            // populateOrders(orderPage);
        }
    })
}

function changePage(page) {
    orderPage = page;
    $("#order_page").val(page)
    shown = false;

    $("#ordersFilterTable tr").filter(function () {
        index = $(this).index()

        if ((index >= ((page - 1) * 25)) && (index < (page * 25))) {
            $(this).toggle(true);
            shown = true;
        }
        else {
            $(this).toggle(false);
        }
    });

    return shown;
}