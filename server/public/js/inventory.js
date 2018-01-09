console.log("I RAN inventory.js");
// $(document).ready( function() {
//     //$("#topnavbar").load("header.html")

//     // inventoryPage = ipcRenderer.sendSync("get_inventory_page", 0);


//     $("#next").click(function () {
//         if (typeof inventoryPage == "string") {
//             inventoryPage = parseInt(inventoryPage, 10);
//         }

//         if (inventoryPage !== null) {
//             // ipcRenderer.send("set_inventory_page", (inventoryPage + 1));
//         }
//     });

//     $("#previous").click(function () {
//         if (typeof inventoryPage == "string") {
//             inventoryPage = parseInt(inventoryPage);
//         }
//         if (inventoryPage !== null && inventoryPage > 1) {
//             ipcRenderer.send("set_inventory_page", inventoryPage - 1);
//         }
//     });

//     $("#gotopagenumber").click(function () {
//         pagenum = $("#pagenumber").val()

//         if (pagenum > 0) {
//             ipcRenderer.send("set_inventory_page", pagenum);
//         }
//     });


//     $.get("http://localhost:9000/v1/inventories/" + 1 + "/20", function (data) {
//         if (data !== null) {
//             for (var i = 0; i < data.length; i++) {
//                 item = data[i];

//                 AddItem(item.inventory_id, item.product.product_name, item.product.product_genre, item.amount, item.product.product_price, item.inventory_condition);
//             }
//         }
//     })
//         .done(function () {
//             //alert("second success");
//         })
//         .fail(function () {
//             //alert("error");
//         })
//         .always(function (data) {
//             //alert("finished" + data);
//         });

// })();

$(document).ready(function () {
    console.log("DOC ready", $.get);

    $.get("http://localhost:9000/v1/inventories/" + 1 + "/20", function (data) {
        if (data !== null) {
            for (var i = 0; i < data.length; i++) {
                item = data[i];

                AddItem(item.inventory_id, item.product.product_name, item.product.product_genre, item.amount, item.product.product_price, item.inventory_condition);
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

console.log("HI:",$(".div"));

var result = "";

function AddItem(itemID, itemName, itemGenre, quantity, price, condition) {
    result = result +
        "<div class='listitem row border border-info'>" +
        // " <a href=' " + item[0] + "'>" + 
        "<div class='col-sm-1'>" + itemID + "</div>" +
        "<div class='col-sm-4'>" + itemName + "</div>" +
        "<div class='col-sm-2'>" + itemGenre + "</div>" +
        "<div class='col-sm-1'>" + quantity + "</div>" +
        "<div class='col-sm-1'>" + price + "</div>" +
        "<div class='col-sm-2'>" + condition + "</div>" +
        "<form class='form-inline my-2 my-lg-0'> <div class='col-sm-1'> <button class='btn-sm btn-outline-success my-2 my-sm-0' onclick='location.href='item.html';'>" + itemID + "</button>" + "</div></form>" +
        // " </a>" +
        "</div>";
    $("#inventory").html(result);
}