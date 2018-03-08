$(document).ready(function() {  

    getBondListFromAPI(0);
    getBondListFromAPI(1);

    /* Using $("body").on to attach click event even on new elements */
    $("body").on("click", "a.bondSellLink", function(e) {
        e.preventDefault();

        $(this).closest("tbody").find("tr").removeClass("table-info");
        $(this).closest("tr").addClass("table-info");
        $("#bondSelected").text(this.text);

        getBondQuotesFromAPI(1, this.text);
        scrollToAnchor('quotes');
    })

    /* Using $("body").on to attach click event even on new elements */
    $("body").on("click", "a.bondBuyLink", function(e) {
        e.preventDefault();

        $(this).closest("tbody").find("tr").removeClass("table-info");
        $(this).closest("tr").addClass("table-info");
        $("#bondSelected").text(this.text);

        getBondQuotesFromAPI(0, this.text);
        scrollToAnchor('quotes');
    });

    /* Create DataTable object */
    $('#tblResult').DataTable({
        "destroy": true
    });
})

function getBondListFromAPI(tp) {
    var data = {};

    data.Kind = tp;

    var json = JSON.stringify(data);                

    /* Send the data using post */
    var posting = $.post("getBondList", json);
 
    /* Put the results in a div */
    posting.done(function(data) {
        var response = JSON.parse(data)

        /* format all dates and times from json before plotting data */
        $.each(response, function(i, bond){
            response[i].DueDate = formatDate(bond.DueDate);
        });

        renderBondListTable(tp, response);
    });

    posting.fail(function(data) {
        alert('Could not retrieve data from server. Is it up and running?');
    });
}

function getBondQuotesFromAPI(tp, name) {
    
    var data = {};

    data.Name = name;
    data.Kind = tp;

    var json = JSON.stringify(data);                

    /* Send the data using post */
    var posting = $.post("getBonds", json);
 
    /* Put the results in a datatable */
    posting.done(function(data) {

        var response = JSON.parse(data);

        /* format all dates and times from json before plotting data */
        $.each(response, function(i, bond){
            response[i].FetchDate = formatDate(bond.FetchDate);
            response[i].DueDate = formatDate(bond.DueDate);
            response[i].LastUpdate = formatDateTime(bond.LastUpdate);
        });

        /* Plot DataTable */
        $('#tblResult').DataTable({
            "destroy": true,
            "data": response,
            "columns": [
                { "data": "FetchDate" },
                { "data": "Rate" },
                { "data": "MinPrice" },
                { "data": "UnitPrice" },
                { "data": "LastUpdate" }
            ],
            "order": [ 0, 'desc' ]
        });
    });

    posting.fail(function(data) {
        alert('Could not retrieve data from server. Is it up and running?');
    });
}   

function renderBondListTable(tp, data) {

    var tagName = "";
    var linkClass = "";

    if (tp == 0) {
        tagName = "table#bondBuyList tbody";
        linkClass = "bondBuyLink";
    } else {
        tagName = "table#bondSellList tbody";
        linkClass = "bondSellLink";
    }

    $(tagName).html("");

    $.each(data, function(i, bond){

        var tblLine = "";

        var tblLine = "<tr>" + 
                "<td><a href='#quotes' class='" + linkClass + "'>" + bond.Name + "</a></td>" +
                "<td>" + bond.Index + "</td>" +
                "<td>" + bond.DueDate + "</td>" +
            "</tr>"

        $(tagName).append(tblLine);
    }); 
}

function formatDate(datetime) {
    var result = "";
    var dt = new Date(datetime);

    month = dt.getMonth() + 1;

    result = dt.getFullYear().toString() + "-" + 
             month.toString().padStart(2,"0") + "-" + 
             dt.getDate().toString().padStart(2,"0");

    return result;
}

function formatDateTime(datetime) {
    var result = "";
    var dt = new Date(datetime);

    month = dt.getMonth() + 1;

    result = dt.getFullYear().toString() + "-" + 
             month.toString().padStart(2,"0") + "-" + 
             dt.getDate().toString().padStart(2,"0") + " " +
             dt.getHours().toString().padStart(2,"0") + ":" + 
             dt.getMinutes().toString().padStart(2,"0") + ":" + 
             dt.getSeconds().toString().padStart(2,"0");

    return result;
}

function scrollToAnchor(aid){
    // var aTag = $("a[name='"+ aid +"']");
    var aTag = $("#bondSelected");
    $('html,body').animate({scrollTop: aTag.offset().top},'slow');
}