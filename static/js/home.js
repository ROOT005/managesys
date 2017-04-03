   $(function(){
    $.ajax({
        url: '/info',
        type: 'get',
        dataType: 'json',
        success:function(data){
            console.log(data);
        },
    });
    
    Chart.defaults.global.defaultFontColor = "#000";
    var data1 = {
        labels: ["星期一", "星期二", "星期三","星期四","星期五","星期六","星期日"],
        datasets:[
            {
                label:"意向客户",
                borderColor: "rgba(168, 176, 225, 1)",
                fill:true,
                lineTension: 0,
                data:[10, 20, 15, 26, 22,10,3]
            },
            {
                label:"签约客户",
                borderColor: "rgba(75, 192, 192, 1)",
                lineTension: 0,
                fill:true,
                data:[8, 15, 20, 17, 20,3,10]
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
        labels: ["测试1","测试2","测试3","测试4","测试5","测试6"],
        datasets:[
            {
                label:"签约客户",
                data:[18,16,20,25,23,17,],
                borderWidth: 1,
                backgroundColor:[
                   "rgba(75, 192, 192, 0.5)",
                    "rgba(75, 192, 192, 0.5)",
                    "rgba(75, 192, 192, 0.5)",
                    "rgba(75, 192, 192, 0.5)",
                    "rgba(75, 192, 192, 0.5)",
                    "rgba(75, 192, 192, 0.5)",
                ],
            },
            {
                label:"意向客户",
                data:[25, 22, 15, 17, 19, 20],
                backgroundColor:[
                    "rgba(255,99,132,0.5)",
                    "rgba(255,99,132,0.5)",
                    "rgba(255,99,132,0.5)",
                    "rgba(255,99,132,0.5)",
                    "rgba(255,99,132,0.5)",
                    "rgba(255,99,132,0.5)",
                ],
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
})