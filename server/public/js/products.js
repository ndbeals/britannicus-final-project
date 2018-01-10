$(document).ready(function () {
    $("#productsFilterInput").on("keyup", function () {
        console.log("GA");
        var value = $(this).val().toLowerCase();
        $("#productsFilterTable tr").filter(function () {
            $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
        });
    });
});