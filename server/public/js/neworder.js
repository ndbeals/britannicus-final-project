inventoryID = 1;
customerID = 1;
itemList ={};

$(document).ready(function () {

    inventoryID = parseInt($("#inventory_ID").val());


    $("#item_next").click(function () {
        inventoryID++;
        populateItem(inventoryID)
    });

    $("#item_previous").click(function () {
        if (inventoryID > 1){
            inventoryID--;

            populateItem(inventoryID)
        }
    });
    
    $("#item_ID").bind('change', function () {
        itemid = parseInt($("#inventory_ID").val());
        
        if (itemid > 0) {
            populateItem(itemid)
        }
    });
    
    
    $("#item_add").click(function () {
        getOrderItem(inventoryID)
    });

    $("#item_clear").click(function () {
        $("#orderlisttable").empty()
        itemList={};
    });

    $("#customer_next").click(function () {
        customerID++;
        populateCustomer(customerID)
    });

    $("#customer_previous").click(function () {
        if (customerID > 1) {
            customerID--;

            populateCustomer(customerID)
        }
    });

    $("#customer_ID").bind('change', function () {
        customerid = parseInt($("#customer_ID").val());

        if (customerid > 0) {
            populateCustomer(customerid)
        }
    });


    $("#order_createform").submit( function(e) {
        e.preventDefault();

        var data = {};
        data.customer_id = customerID;
        data.item_list = itemList;

        $.ajax({
            url: "/v1/order",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                // console.log("DATA POSTED SUCCESSFULLY" , data);
                window.location.href="/order/get/"+data.data.order_id
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });

    // $("#inventory_condition").val(inventoryCondtion)

    populateItem(1)
    populateCustomer(1)
});

function populateItem( id ) {
    $.get("/v1/inventory/"+id, function(data) {
        if (data !== null) {
            inventoryID = id
            $("#item_ID").val(id)

            $("#item_isbn").val(data.product.isbn)
            $("#item_author").val(data.product.product_author)
            $("#item_genre").val(data.product.product_genre)
            $("#item_publisher").val(data.product.product_publisher)
            $("#item_description").val(data.product.product_description)
            $("#item_name").val(data.product.product_name)
            $("#item_type").val(data.product.product_type)
            $("#inventory_price").val(data.item_price)
        }
    })
}

function populateCustomer(id) {
    $.get("/v1/customer/" + id, function (data) {
        if (data !== null) {
            customerID = id
            $("#customer_ID").val(id)

            $("#customer_firstname").val( data.first_name ) 
            $("#customer_lastname").val( data.last_name ) 
            $("#customer_email").val( data.customer_email ) 
            $("#customer_phonenumber").val( data.customer_phone ) 
            $("#customer_address").val( data.customer_address ) 
            $("#customer_city").val( data.customer_city ) 
            $("#customer_state").val( data.customer_state ) 
            $("#customer_country").val( data.customer_country ) 
        }
    })
}

function updateList() {
    $.get("/v1/order/" + orderID, function (data) {
        if (data !== null) {
            table = $("#orderlisttable")

            for (var i = 0; i < data.item_list.length; i++) {

                addOrderItem(table, orderID, data.item_list[i])
            }
        }
    })
}

function getOrderItem( id ) {
    $.get("/v1/inventory/" + id, function (data) {
        if (data !== null) {
            amount = parseInt($("#order_amount").val())

            if (amount > 0 && amount < 100) {
                table = $("#orderlisttable")

                addOrderItem( table ,amount, data )
                
            }
        }
    })
}

function addOrderItem(table, amount, item) {

    idint = parseInt(item.inventory_id)
    itemList[idint]=amount

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
                `+ amount + `
            </div>
        </td>
        <td>
            <div>
                `+ ( parseFloat(item.item_price) * amount) + `
            </div>
        </td>
    </tr>`

    row = $(row).appendTo(table)

    return row;
}