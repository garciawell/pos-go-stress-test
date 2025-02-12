package web

const HtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Resultados do Load Test</title>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100..900;1,100..900&display=swap" rel="stylesheet">
    <style>
        body{
            font-family: 'Roboto', sans-serif;
        }
    </style>
</head>
<body>
    <h2>Resultados do Load Test</h2>
    <canvas id="chart"></canvas>
    <script>
        const data = {{ . }};
        const ctx = document.getElementById('chart').getContext('2d');
        new Chart(ctx, {
            type: 'bar',
            data: {
                labels: ['Requests', 'Média Latência (ms)', 'Sucesso (%)'],
                datasets: [{
                    label: 'Métricas',
                    data: [data.requests, data.mean_latency_ms, data.success_rate, data.error_rate],
                    backgroundColor: ['blue', 'yellow', 'green', 'red']
                }]
            }
        });
    </script>
</body>
</html>
`
