let growthChart = null;

document.getElementById('calc-form').addEventListener('submit', async (e) => {
    e.preventDefault();

    const rawRate = parseFloat(document.getElementById('rate').value);

    const formData = {
        principal: parseFloat(document.getElementById('principal').value),
        rate: rawRate / 100, // convert percent to decimal
        compoundsPerYear: parseInt(document.getElementById('compoundsPerYear').value),
        years: parseFloat(document.getElementById('years').value),
    };

    if (Object.values(formData).some(v => isNaN(v) || v <= 0)) {
        document.getElementById('result').innerText = 'Please enter valid positive numbers.';
        return;
    }

    try {
        const response = await fetch('/api/v1/calculate', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(formData),
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        document.getElementById('result').innerText = `Result is: $${data.result.toFixed(2)}`;

        const chartData = [];
        for (let year = 0; year <= formData.years; year++) {
            const amount = formData.principal * Math.pow(
                1 + formData.rate / formData.compoundsPerYear,
                formData.compoundsPerYear * year
            );
            chartData.push(amount);
        }

        const ctx = document.getElementById('growthChart').getContext('2d');

        if (growthChart) {
            growthChart.destroy();
        }

        growthChart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: Array.from({ length: chartData.length }, (_, i) => `Year ${i}`),
                datasets: [{
                    label: 'Investment Growth',
                    data: chartData,
                    borderColor: 'rgba(59, 130, 246, 1)',
                    backgroundColor: 'rgba(59, 130, 246, 0.2)',
                    fill: true,
                }],
            },
            options: {
                responsive: true,
                plugins: {
                    legend: { position: 'top' },
                    title: { display: true, text: 'Compound Interest Growth' },
                },
                scales: {
                    y: { title: { display: true, text: 'Amount ($)' } },
                    x: { title: { display: true, text: 'Years' } },
                },
            },
        });
    } catch (error) {
        document.getElementById('result').innerText = `Error: ${error.message}`;
    }
});
