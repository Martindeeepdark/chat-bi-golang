document.addEventListener('DOMContentLoaded', function () {
    // 获取 echarts 容器
    var chartContainer = document.getElementById('echart');

    // 初始化 echarts 实例
    var myChart = echarts.init(chartContainer);
    // 使用 fetch API 发送 POST 请求到 /echat 路由
    fetch('/echat', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestData)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            // 检查返回的数据格式
            if (data && data.echat) {
                // 使用返回的 echarts 数据配置图表
                myChart.setOption(data.echat);
            } else {
                console.error('Invalid data format:', data);
            }
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
});