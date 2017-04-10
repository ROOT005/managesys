   $(function(){
    $.ajax({
        url: '/info',
        type: 'get',
        dataType: 'json',
        success:function(data){
            var info1 = data.weekInfo;
            var weekinfo = new Array();
            for (var i = 0; i < 7; i++) {
                weekinfo[i] = info1[i];
            }

            var info2 = data.weekInfoFin;
            var weekinfofin = new Array();
            for (var i = 0; i < 7; i++) {
                weekinfofin[i] = info2[i];
            }
            console.log(weekinfofin)
            console.log(weekinfo)
            var info3 = data.dayInfo;
            var dayinfouser = new Array();
            var dayinfovalue = new Array();
            for(var key in info3){
                dayinfouser.push(key); 
                dayinfovalue.push(info3[key]);
            }
            var info4 = data.dayInfoFin;
            var dayinfofinvalue = new Array();
            for(var key in info3){
                dayinfofinvalue.push(info4[key]);
            }
        
            Chart.defaults.global.defaultFontColor = "#4bc0de";
            var data1 = {
                labels: ["星期一", "星期二", "星期三","星期四","星期五","星期六","星期日"],
                datasets:[
                    {
                        label:"意向客户",
                        borderColor: "rgba(168, 176, 225, 1)",
                        fill:true,
                        lineTension: 0,
                        data: weekinfo, //[10, 20, 15, 26, 22,10,3] //
                    },
                    {
                        label:"签约客户",
                        borderColor: "rgba(75, 192, 192, 1)",
                        lineTension: 0,
                        fill:true,
                        data:weekinfofin,//[8, 15, 20, 17, 20,3,10] 
                    }
                ]
            };
            var options1 = {
            }
            var ctx1 = $("#chart1").get(0).getContext("2d");
            var myChart1 = Chart.Line(ctx1, {
                data: data1,
                options: options1,
            });

            //每日报表
            var data2 = {
                labels: dayinfouser,
                datasets:[
                    {
                        label:"签约客户",
                        data:dayinfofinvalue,
                        borderWidth: 1,
                        backgroundColor:"rgba(75, 192, 192, 0.5)",
                    },
                    {
                        label:"意向客户",
                        data:dayinfovalue,
                        backgroundColor:"rgba(255,99,132,0.5)",
                    }
                ],

            }
            var options2 = {
                
            }
            var ctx2 = $("#chart2").get(0).getContext("2d");
            var myChart2 = new Chart(ctx2, {
                type: "bar",
                data: data2,
                options: options2,
            })
        },
    });

    
})
