inventoryPage = 1;

$(document).ready(function () {
    inventoryPage = parseInt($("#inventory_page").val());


    $("#inventory_next").click(function () {
        if (!changePage(inventoryPage + 1)) {
            changePage(inventoryPage - 1);
        }
    });

    $("#inventory_previous").click(function () {
        if (!changePage(inventoryPage - 1)) {
            changePage(inventoryPage + 1);
        }
    });

    $("#inventory_page").bind('keyup input', function () {
        pagenum = parseInt($("#inventory_page").val());

        if (pagenum > 0) {
            oldpage = inventoryPage;

            if (!changePage(pagenum)) {
                changePage(oldpage);
            }
        }
    });


    $("#inventorysFilterInput").on("keyup", function () {
        var value = $(this).val().toLowerCase();

        if (value == "") {
            changePage(inventoryPage);
        } else {
            results = 0
            $("#inventorysFilterTable tr").filter(function () {
                if ($(this).is(":visible")) {
                    $(this).toggle(false)
                }
                if (results < 15) {
                    show = $(this).text().toLowerCase().indexOf(value) > -1
                    $(this).toggle(show)
                    if (show) {
                        results++;
                    }
                }
            });
        }
    });

    populateInventory(1);
});

function addInventory(table, inventory_id, product_name, product_author, product_genre, amount, price, inventory_condition, note) {

    var row = "<tr><td>" +
        inventory_id + "</td><td>" +
        product_name + "</td><td>" +
        product_author + "</td><td>" +
        product_genre + "</td><td>" +
        amount + "</td><td>$" +
        price + "</td><td>" +
        inventory_condition + "</td><td>" +
        note + '</td><td><a href="/inventory/get/' + inventory_id + '"><button type="button" class="btn btn-primary btn-block tbl-btn">Edit</button></a></td><td><button id="btndel' + inventory_id + '" type="button" class="btn btn-danger btn-block tbl-btn">Delete</button></td></tr>'

    row = $(row).appendTo(table)//table.append(row);

    $('#btndel' + inventory_id).click(function () {
        var parent = $(this).parent().parent();
        $.ajax({
            url: "/v1/inventory/" + inventory_id,
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

function populateInventory(page, hide) {
    $.get("/v1/inventories/" + page + "/15", function (data) {
        if (data !== null) {
            var table = $("#inventorysFilterTable");
            
            // table.empty();
            
            for (var i = 0; i < data.length; i++) {
                item = data[i];
                // console.log(item)

                row = addInventory(table, item.inventory_id, item.product.product_name, item.product.product_author, item.product.product_genre, item.amount, item.item_price, item.inventory_condition, item.note)

                if (hide == true) {
                    row.toggle(false)
                    // changePage(inventoryPage);
                }
            }

            populateInventory(page + 1, true);
        }
        else {
            // populateInventorys(inventoryPage);
        }
    })
        .done(function () {
            //alert("second success");
        })
        .fail(function (data) {
            //alert("error");
        })
        .always(function (data) {
            //alert("finished" + data);
        });
}

function changePage(page) {
    inventoryPage = page;
    $("#inventory_page").val(page)
    shown = false;

    $("#inventorysFilterTable tr").filter(function () {
        index = $(this).index()

        if ((index >= ((page - 1) * 15)) && (index < (page * 15))) {
            $(this).toggle(true);
            shown = true;
        }
        else {
            $(this).toggle(false);
        }
    });

    return shown;
}