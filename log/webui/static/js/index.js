async function loadLevels() {
    const res = await fetch("/log/level/all");
    return await res.json();
}

async function loadPackages() {
    const res = await fetch("/log/level/package/all");
    return await res.json();
}

async function render() {
    const levels = await loadLevels();
    const pkgs = await loadPackages();

    const tbody = document.getElementById("pkgTableBody");
    tbody.innerHTML = "";

    pkgs.forEach(pkg => {
        const tr = document.createElement("tr");

        // Package column
        const tdName = document.createElement("td");
        tdName.textContent = pkg.name;

        // Level dropdown
        const tdLevel = document.createElement("td");
        const select = document.createElement("select");

        for (const lvl in levels) {
            const opt = document.createElement("option");
            opt.value = lvl;
            opt.textContent = levels[lvl];
            if (parseInt(lvl) === pkg.level) opt.selected = true;
            select.appendChild(opt);
        }

        tdLevel.appendChild(select);

        // Save button
        const tdAction = document.createElement("td");
        const button = document.createElement("button");
        button.textContent = "Save";

        button.onclick = async () => {
            const params = new URLSearchParams();
            params.append("pkg", pkg.name);
            params.append("level", select.value);

            await fetch("/log/level/package/update", {
                method: "POST",
                headers: { "Content-Type": "application/x-www-form-urlencoded" },
                body: params.toString()
            });

            await render();
        };

        tdAction.appendChild(button);

        // Assemble row
        tr.appendChild(tdName);
        tr.appendChild(tdLevel);
        tr.appendChild(tdAction);

        tbody.appendChild(tr);
    });
}

render();