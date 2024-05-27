document.addEventListener('DOMContentLoaded', function() {
    const uploadButton = document.getElementById('uploadButton');
    if (uploadButton) {
        uploadButton.addEventListener('click', uploadFile);
    }
});

function uploadFile() {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];
    const request = document.getElementById('analysisRequest').value;

    if (!file) {
        alert('请先选择一个文件');
        return;
    }

    const formData = new FormData();
    formData.append('file', file);
    formData.append('request', request);

    const statusDiv = document.getElementById('status');
    const analysisReport = document.getElementById('analysisReport');

    statusDiv.style.display = 'block';
    statusDiv.innerText = '正在上传...';
    statusDiv.style.backgroundColor = '#ddd';

    fetch('/process', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('网络响应不正常');
            }
            return response.json();
        })
        .then(data => {
            statusDiv.innerText = '上传成功: ' + data.message;
            console.log(statusDiv.innerText)
            console.log(data.message)
            console.log(data.analysis_report)
            statusDiv.style.backgroundColor = '#c8e6c9'; // 浅绿色

            // 显示分析结果
            analysisReport.value = data.analysis_report;
        })
        .catch(error => {
            statusDiv.innerText = '上传失败: ' + error.message;
            statusDiv.style.backgroundColor = '#ffccbc'; // 浅红色
        });
}