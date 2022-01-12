module.exports.timestamptodate = function(timestamp){
    const today = new Date(timestamp);
    const date = today.getDate()+'/'+(today.getMonth()+1)+'/'+today.getFullYear()
    const time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
    const dateTime = date +' '+ time;
    timestamp = dateTime;
    return timestamp;
}
module.exports.datetotimestamp = function(myDate){
    //var myDate = "26-02-2012";
    myDate = myDate.split("/");
    var newDate = new Date( myDate[2], myDate[1] - 1, myDate[0]);
    //console.log(newDate.getTime());
    return newDate.getTime();
}
