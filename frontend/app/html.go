package app

// EmbeddedHTML contains the complete HTML content for the wallet generator frontend
const EmbeddedHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>SamWallet</title>
        <style>
            * {
                margin: 0;
                padding: 0;
                box-sizing: border-box;
            }

            body {
                font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
                max-width: 1400px;
                margin: 0 auto;
                padding: 20px;
                background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
                min-height: 100vh;
                color: #333;
            }

            .container {
                display: flex;
                gap: 24px;
                margin-top: 20px;
            }

            .panel {
                background: rgba(255, 255, 255, 0.95);
                backdrop-filter: blur(10px);
                padding: 32px;
                border-radius: 20px;
                box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
                flex: 1;
                border: 1px solid rgba(255, 255, 255, 0.2);
                transition: all 0.3s ease;
            }

            .panel:hover {
                transform: translateY(-2px);
                box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2);
            }

            .title {
                text-align: center;
                background: linear-gradient(135deg, #667eea, #764ba2);
                -webkit-background-clip: text;
                -webkit-text-fill-color: transparent;
                background-clip: text;
                margin-bottom: 40px;
                font-size: 32px;
                font-weight: 700;
                letter-spacing: -1px;
            }

            .section {
                margin-bottom: 28px;
            }

            .section h3 {
                margin-top: 0;
                margin-bottom: 16px;
                color: #2c3e50;
                font-size: 18px;
                font-weight: 600;
                position: relative;
                padding-bottom: 8px;
            }

            .section h3::after {
                content: "";
                position: absolute;
                bottom: 0;
                left: 0;
                width: 40px;
                height: 3px;
                background: linear-gradient(135deg, #667eea, #764ba2);
                border-radius: 2px;
            }

            input,
            select {
                width: 100%;
                padding: 14px 16px;
                margin: 8px 0;
                border: 2px solid #e1e8ed;
                border-radius: 12px;
                font-size: 15px;
                transition: all 0.3s ease;
                background: rgba(255, 255, 255, 0.8);
            }

            input:focus,
            select:focus {
                outline: none;
                border-color: #667eea;
                background: white;
                box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1);
                transform: translateY(-1px);
            }

            button {
                width: 100%;
                padding: 14px 20px;
                margin: 8px 0;
                border: none;
                border-radius: 12px;
                font-size: 15px;
                font-weight: 600;
                cursor: pointer;
                transition: all 0.3s ease;
                position: relative;
                overflow: hidden;
                background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
                color: white;
                box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
            }

            button::before {
                content: "";
                position: absolute;
                top: 0;
                left: -100%;
                width: 100%;
                height: 100%;
                background: linear-gradient(
                    90deg,
                    transparent,
                    rgba(255, 255, 255, 0.2),
                    transparent
                );
                transition: left 0.5s;
            }

            button:hover::before {
                left: 100%;
            }

            button:hover {
                transform: translateY(-2px);
                box-shadow: 0 12px 30px rgba(102, 126, 234, 0.4);
            }

            button:active {
                transform: translateY(0);
            }

            button.delete {
                background: linear-gradient(135deg, #ff6b6b 0%, #ee5a52 100%);
                box-shadow: 0 8px 20px rgba(255, 107, 107, 0.3);
            }

            button.delete:hover {
                box-shadow: 0 12px 30px rgba(255, 107, 107, 0.4);
            }
            .list {
                max-height: 320px;
                overflow-y: auto;
                border: 2px solid #e1e8ed;
                border-radius: 16px;
                padding: 8px;
                background: rgba(255, 255, 255, 0.5);
                backdrop-filter: blur(5px);
            }

            .list::-webkit-scrollbar {
                width: 6px;
            }

            .list::-webkit-scrollbar-track {
                background: rgba(0, 0, 0, 0.1);
                border-radius: 3px;
            }

            .list::-webkit-scrollbar-thumb {
                background: linear-gradient(135deg, #667eea, #764ba2);
                border-radius: 3px;
            }

            .list-item {
                padding: 14px 16px;
                margin: 4px 0;
                border-radius: 12px;
                cursor: pointer;
                transition: all 0.3s ease;
                position: relative;
                border: 1px solid transparent;
                font-weight: 500;
            }

            .list-item:hover {
                background: linear-gradient(
                    135deg,
                    rgba(102, 126, 234, 0.1),
                    rgba(118, 75, 162, 0.1)
                );
                border-color: rgba(102, 126, 234, 0.3);
                transform: translateX(4px);
            }

            .list-item.selected {
                background: linear-gradient(
                    135deg,
                    rgba(102, 126, 234, 0.15),
                    rgba(118, 75, 162, 0.15)
                );
                border-color: #667eea;
                color: #2c3e50;
                font-weight: 600;
                box-shadow: 0 4px 12px rgba(102, 126, 234, 0.2);
            }

            .status {
                text-align: center;
                padding: 16px 20px;
                margin-top: 24px;
                border-radius: 16px;
                font-weight: 600;
                position: relative;
                overflow: hidden;
                border: none;
            }

            .status::before {
                content: "";
                position: absolute;
                top: 0;
                left: 0;
                right: 0;
                height: 4px;
                background: currentColor;
                opacity: 0.3;
            }

            .status.success {
                background: linear-gradient(135deg, #10b981, #059669);
                color: white;
                box-shadow: 0 8px 20px rgba(16, 185, 129, 0.3);
            }

            .status.error {
                background: linear-gradient(135deg, #ef4444, #dc2626);
                color: white;
                box-shadow: 0 8px 20px rgba(239, 68, 68, 0.3);
            }

            .status.info {
                background: linear-gradient(135deg, #3b82f6, #1d4ed8);
                color: white;
                box-shadow: 0 8px 20px rgba(59, 130, 246, 0.3);
            }

            /* Modern Modal Styles */
            .modal {
                display: none;
                position: fixed;
                z-index: 1000;
                left: 0;
                top: 0;
                width: 100%;
                height: 100%;
                background: rgba(0, 0, 0, 0.7);
                backdrop-filter: blur(8px);
                animation: modalFadeIn 0.3s ease;
            }

            @keyframes modalFadeIn {
                from {
                    opacity: 0;
                }
                to {
                    opacity: 1;
                }
            }

            @keyframes modalSlideIn {
                from {
                    transform: translateY(-50px);
                    opacity: 0;
                }
                to {
                    transform: translateY(0);
                    opacity: 1;
                }
            }

            .modal-content {
                background: rgba(255, 255, 255, 0.95);
                backdrop-filter: blur(20px);
                margin: 3% auto;
                padding: 0;
                border: 1px solid rgba(255, 255, 255, 0.3);
                width: 90%;
                max-width: 650px;
                border-radius: 24px;
                box-shadow: 0 25px 50px rgba(0, 0, 0, 0.3);
                position: relative;
                animation: modalSlideIn 0.4s ease;
                overflow: hidden;
            }

            .modal-header {
                background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
                color: white;
                padding: 24px 32px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                position: relative;
            }

            .modal-header::before {
                content: "";
                position: absolute;
                top: 0;
                left: 0;
                right: 0;
                bottom: 0;
                background: linear-gradient(
                    135deg,
                    rgba(255, 255, 255, 0.1) 0%,
                    transparent 100%
                );
            }

            .modal-header h3 {
                margin: 0;
                font-size: 20px;
                font-weight: 700;
                letter-spacing: -0.5px;
                position: relative;
                z-index: 1;
            }

            .close {
                color: white;
                font-size: 28px;
                font-weight: bold;
                cursor: pointer;
                background: rgba(255, 255, 255, 0.1);
                border: 2px solid rgba(255, 255, 255, 0.2);
                padding: 8px 12px;
                width: auto;
                border-radius: 50%;
                transition: all 0.3s ease;
                position: relative;
                z-index: 1;
            }

            .close:hover {
                background: rgba(255, 255, 255, 0.2);
                transform: rotate(90deg);
            }

            .modal-body {
                padding: 32px;
            }

            .modal-footer {
                padding: 24px 32px;
                background: rgba(0, 0, 0, 0.02);
                border-top: 1px solid rgba(0, 0, 0, 0.1);
                text-align: right;
            }

            .warning {
                color: white;
                font-weight: 600;
                background: linear-gradient(135deg, #ff6b6b, #ee5a52);
                border: none;
                padding: 16px 20px;
                border-radius: 16px;
                margin: 16px 0;
                box-shadow: 0 8px 20px rgba(255, 107, 107, 0.3);
                position: relative;
                overflow: hidden;
            }

            .warning::after {
                content: "";
                position: absolute;
                top: 0;
                left: 0;
                right: 0;
                height: 4px;
                background: rgba(255, 255, 255, 0.3);
            }

            .info {
                color: white;
                font-weight: 600;
                background: linear-gradient(135deg, #10b981, #059669);
                border: none;
                padding: 16px 20px;
                border-radius: 16px;
                margin: 16px 0;
                box-shadow: 0 8px 20px rgba(16, 185, 129, 0.3);
                position: relative;
                overflow: hidden;
            }

            .info::after {
                content: "";
                position: absolute;
                top: 0;
                left: 0;
                right: 0;
                height: 4px;
                background: rgba(255, 255, 255, 0.3);
            }

            /* Mnemonic container styles */
            .mnemonic-container {
                margin: 15px 0;
            }

            .mnemonic-input-group {
                position: relative;
                display: block;
            }

            .mnemonic-input-group textarea {
                width: 100%;
                padding-right: 50px !important;
                box-sizing: border-box;
                resize: vertical;
                min-height: 80px;
                max-width: 100%;
            }

            .eye-button {
                position: absolute;
                right: 12px;
                top: 12px;
                background: rgba(255, 255, 255, 0.95);
                backdrop-filter: blur(10px);
                border: 2px solid rgba(102, 126, 234, 0.2);
                font-size: 16px;
                cursor: pointer;
                padding: 8px 10px;
                border-radius: 12px;
                transition: all 0.3s ease;
                z-index: 10;
                width: auto;
                height: auto;
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            }

            .eye-button:hover {
                background: linear-gradient(
                    135deg,
                    rgba(102, 126, 234, 0.1),
                    rgba(118, 75, 162, 0.1)
                );
                border-color: #667eea;
                transform: scale(1.05);
                box-shadow: 0 6px 16px rgba(102, 126, 234, 0.2);
            }

            .mnemonic-hidden {
                -webkit-text-security: disc;
                text-security: disc;
            }

            /* Responsive Design */
            @media (max-width: 768px) {
                .container {
                    flex-direction: column;
                    gap: 16px;
                }

                .panel {
                    padding: 24px;
                    border-radius: 16px;
                }

                .title {
                    font-size: 28px;
                    margin-bottom: 32px;
                }

                .modal-content {
                    width: 95%;
                    margin: 5% auto;
                }

                .modal-body {
                    padding: 24px;
                }

                .modal-header {
                    padding: 20px 24px;
                }

                .modal-footer {
                    padding: 20px 24px;
                }
            }

            /* Loading animations */
            @keyframes pulse {
                0% {
                    opacity: 1;
                }
                50% {
                    opacity: 0.5;
                }
                100% {
                    opacity: 1;
                }
            }

            .loading {
                animation: pulse 1.5s infinite;
            }

            /* Scroll animations */
            @keyframes slideInUp {
                from {
                    transform: translateY(30px);
                    opacity: 0;
                }
                to {
                    transform: translateY(0);
                    opacity: 1;
                }
            }

            .animate-in {
                animation: slideInUp 0.6s ease forwards;
            }

            .title h4 {
                color: white !important;
                -webkit-text-fill-color: white;
            }
        </style>
    </head>
    <body>
        <div class="title">
            <h4>SamWallet - Crypto Wallet Generator (BTC & ETH)</h4>
        </div>

        <div class="container">
            <div class="panel animate-in">
                <div class="section">
                    <h3>🔐 Create New Wallet</h3>
                    <input
                        type="text"
                        id="walletName"
                        placeholder="Enter wallet name..."
                    />
                    <button onclick="createWallet()">Create Wallet</button>
                </div>

                <div class="section">
                    <h3>💼 Existing Wallets</h3>
                    <div id="walletList" class="list"></div>
                    <button class="delete" onclick="deleteWallet()">
                        Delete Selected Wallet
                    </button>
                </div>
            </div>

            <div class="panel animate-in">
                <div class="section">
                    <h3>🚀 Generate Addresses</h3>
                    <select id="coinSelect">
                        <option value="btc">Bitcoin (BTC)</option>
                        <option value="eth">Ethereum (ETH)</option>
                    </select>
                    <button onclick="generateAddress()">
                        Generate Address
                    </button>
                </div>

                <div class="section">
                    <h3>📍 Wallet Addresses</h3>
                    <div id="addressList" class="list"></div>
                </div>
            </div>
        </div>

        <div id="status" class="status"></div>

        <!-- Mnemonic Modal -->
        <div id="mnemonicModal" class="modal" style="display: none">
            <div class="modal-content">
                <div class="modal-header">
                    <h3>Wallet Created Successfully!</h3>
                    <span class="close" onclick="closeMnemonicModal()"
                        >&times;</span
                    >
                </div>
                <div class="modal-body">
                    <p>
                        <strong>Wallet Name:</strong>
                        <span id="modalWalletName"></span>
                    </p>
                    <p><strong>Your 12-word mnemonic phrase:</strong></p>
                    <p class="warning">
                        ⚠️ Keep this safe! Anyone with these words can access
                        your wallet!
                    </p>
                    <p class="warning">
                        ⚠️ This will be shown only once! Make sure to save it
                        securely!
                    </p>

                    <div class="mnemonic-container">
                        <div class="mnemonic-input-group">
                            <textarea
                                id="mnemonicText"
                                readonly
                                rows="3"
                                style="
                                    font-family: monospace;
                                    font-size: 14px;
                                    padding: 10px;
                                    border: 2px solid #007bff;
                                    border-radius: 4px;
                                    background-color: #f8f9fa;
                                "
                            ></textarea>
                            <button
                                type="button"
                                id="toggleMnemonic"
                                class="eye-button"
                                onclick="toggleMnemonicVisibility()"
                            >
                                <span id="eyeIcon">👁️</span>
                            </button>
                        </div>
                    </div>

                    <p class="info">
                        ✅ Mnemonic has been copied to your clipboard
                    </p>
                </div>
                <div class="modal-footer">
                    <button
                        onclick="closeMnemonicModal()"
                        style="background-color: #28a745"
                    >
                        I've Saved My Mnemonic
                    </button>
                </div>
            </div>
        </div>

        <!-- Delete Wallet Modal -->
        <div id="deleteModal" class="modal" style="display: none">
            <div class="modal-content">
                <div class="modal-header" style="background-color: #dc3545">
                    <h3>Delete Wallet</h3>
                    <span class="close" onclick="closeDeleteModal()"
                        >&times;</span
                    >
                </div>
                <div class="modal-body">
                    <p>
                        <strong
                            >Are you sure you want to delete this
                            wallet?</strong
                        >
                    </p>
                    <p>
                        <strong>Wallet Name:</strong>
                        <span id="deleteWalletName"></span>
                    </p>
                    <p class="warning">⚠️ This action cannot be undone!</p>
                    <p class="warning">
                        ⚠️ All addresses associated with this wallet will be
                        permanently deleted!
                    </p>
                    <p class="warning">
                        ⚠️ Make sure you have backed up your mnemonic phrase!
                    </p>
                </div>
                <div class="modal-footer">
                    <button
                        onclick="closeDeleteModal()"
                        style="background-color: #6c757d; margin-right: 10px"
                    >
                        Cancel
                    </button>
                    <button
                        onclick="confirmDeleteWallet()"
                        style="background-color: #dc3545"
                    >
                        Delete Wallet
                    </button>
                </div>
            </div>
        </div>

        <script>
            let selectedWallet = "";

            // Load wallets on page load
            window.onload = function () {
                loadWallets();
            };

            function showStatus(message, type) {
                const status = document.getElementById("status");
                status.textContent = message;
                status.className = "status " + type;
                setTimeout(() => {
                    status.textContent = "";
                    status.className = "status";
                }, 5000);
            }

            function loadWallets() {
                fetch("/api/wallets")
                    .then((response) => response.json())
                    .then((data) => {
                        if (data.success) {
                            const walletList =
                                document.getElementById("walletList");
                            walletList.innerHTML = "";

                            // Check if data.data exists and is an array
                            if (data.data && Array.isArray(data.data)) {
                                data.data.forEach((wallet) => {
                                    const item = document.createElement("div");
                                    item.className = "list-item";
                                    item.textContent =
                                        wallet.name +
                                        " (Created: " +
                                        wallet.created_at +
                                        ")";
                                    item.onclick = () =>
                                        selectWallet(wallet.name, item);
                                    walletList.appendChild(item);
                                });
                                showStatus(
                                    "Loaded " + data.data.length + " wallets",
                                    "info"
                                );
                            } else {
                                // No wallets or data is null
                                showStatus("No wallets found", "info");
                            }
                        } else {
                            showStatus(
                                "Error loading wallets: " + data.message,
                                "error"
                            );
                        }
                    })
                    .catch((error) => {
                        showStatus("Error loading wallets: " + error, "error");
                    });
            }

            function selectWallet(name, element) {
                selectedWallet = name;
                // Remove previous selection
                document.querySelectorAll(".list-item").forEach((item) => {
                    item.classList.remove("selected");
                });
                // Add selection to clicked item
                element.classList.add("selected");
                loadAddresses(name);
            }

            function loadAddresses(walletName) {
                fetch("/api/addresses?wallet=" + encodeURIComponent(walletName))
                    .then((response) => response.json())
                    .then((data) => {
                        const addressList =
                            document.getElementById("addressList");
                        addressList.innerHTML = "";
                        if (
                            data.success &&
                            data.data &&
                            Array.isArray(data.data) &&
                            data.data.length > 0
                        ) {
                            data.data.forEach((addr, index) => {
                                const item = document.createElement("div");
                                item.className = "list-item";
                                item.textContent =
                                    index +
                                    1 +
                                    ". " +
                                    addr.coin.toUpperCase() +
                                    ": " +
                                    addr.address;
                                addressList.appendChild(item);
                            });
                            showStatus(
                                "Loaded " + data.data.length + " addresses",
                                "info"
                            );
                        } else {
                            showStatus(
                                "No addresses found for this wallet",
                                "info"
                            );
                        }
                    })
                    .catch((error) => {
                        showStatus(
                            "Error loading addresses: " + error,
                            "error"
                        );
                    });
            }

            function createWallet() {
                const name = document.getElementById("walletName").value;
                if (!name) {
                    showStatus("Please enter a wallet name", "error");
                    return;
                }

                fetch("/api/create-wallet", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ name: name }),
                })
                    .then((response) => response.json())
                    .then((data) => {
                        if (data.success) {
                            showStatus(
                                "Wallet created successfully!",
                                "success"
                            );
                            document.getElementById("walletName").value = "";
                            loadWallets();

                            // Show mnemonic modal
                            if (data.data && data.data.mnemonic) {
                                showMnemonicModal(
                                    data.data.wallet_name,
                                    data.data.mnemonic
                                );
                            }
                        } else {
                            showStatus(
                                "Error creating wallet: " + data.message,
                                "error"
                            );
                        }
                    })
                    .catch((error) => {
                        showStatus("Error creating wallet: " + error, "error");
                    });
            }

            function showMnemonicModal(walletName, mnemonic) {
                document.getElementById("modalWalletName").textContent =
                    walletName;
                document.getElementById("mnemonicText").value = mnemonic;
                document.getElementById("mnemonicModal").style.display =
                    "block";

                // Start with mnemonic hidden by default for security
                const textarea = document.getElementById("mnemonicText");
                const eyeIcon = document.getElementById("eyeIcon");
                textarea.classList.add("mnemonic-hidden");
                eyeIcon.textContent = "👁️‍🗨️";
                eyeIcon.title = "Show mnemonic";

                // Copy to clipboard
                navigator.clipboard
                    .writeText(mnemonic)
                    .then(() => {
                        console.log("Mnemonic copied to clipboard");
                    })
                    .catch((err) => {
                        console.error("Failed to copy mnemonic: ", err);
                    });

                // Auto-hide mnemonic after 30 seconds for security (if it was shown)
                setTimeout(() => {
                    if (!textarea.classList.contains("mnemonic-hidden")) {
                        textarea.classList.add("mnemonic-hidden");
                        eyeIcon.textContent = "👁️‍🗨️";
                        eyeIcon.title = "Show mnemonic";
                    }
                }, 30000); // 30 seconds
            }

            function toggleMnemonicVisibility() {
                const textarea = document.getElementById("mnemonicText");
                const eyeIcon = document.getElementById("eyeIcon");

                if (textarea.classList.contains("mnemonic-hidden")) {
                    // Show mnemonic
                    textarea.classList.remove("mnemonic-hidden");
                    eyeIcon.textContent = "👁️";
                    eyeIcon.title = "Hide mnemonic";
                } else {
                    // Hide mnemonic
                    textarea.classList.add("mnemonic-hidden");
                    eyeIcon.textContent = "👁️‍🗨️";
                    eyeIcon.title = "Show mnemonic";
                }
            }

            function closeMnemonicModal() {
                document.getElementById("mnemonicModal").style.display = "none";
            }

            // Modal can only be closed by clicking the close button or confirmation button
            // Clicking outside does not close the modal for security

            function deleteWallet() {
                if (!selectedWallet) {
                    showStatus("Please select a wallet to delete", "error");
                    return;
                }

                // Show delete confirmation modal
                showDeleteModal(selectedWallet);
            }

            function showDeleteModal(walletName) {
                document.getElementById("deleteWalletName").textContent =
                    walletName;
                document.getElementById("deleteModal").style.display = "block";
            }

            function closeDeleteModal() {
                document.getElementById("deleteModal").style.display = "none";
            }

            function confirmDeleteWallet() {
                const walletName =
                    document.getElementById("deleteWalletName").textContent;

                fetch("/api/delete-wallet", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ name: walletName }),
                })
                    .then((response) => response.json())
                    .then((data) => {
                        if (data.success) {
                            showStatus(
                                "Wallet deleted successfully!",
                                "success"
                            );
                            selectedWallet = "";
                            document.getElementById("addressList").innerHTML =
                                "";
                            loadWallets();
                            closeDeleteModal();
                        } else {
                            showStatus(
                                "Error deleting wallet: " + data.message,
                                "error"
                            );
                        }
                    })
                    .catch((error) => {
                        showStatus("Error deleting wallet: " + error, "error");
                    });
            }

            function generateAddress() {
                if (!selectedWallet) {
                    showStatus("Please select a wallet first", "error");
                    return;
                }

                const coin = document.getElementById("coinSelect").value;

                fetch("/api/generate-address", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        wallet: selectedWallet,
                        coin: coin,
                    }),
                })
                    .then((response) => response.json())
                    .then((data) => {
                        if (data.success) {
                            showStatus(
                                coin.toUpperCase() +
                                    " address generated successfully!",
                                "success"
                            );
                            loadAddresses(selectedWallet);
                        } else {
                            showStatus(
                                "Error generating address: " + data.message,
                                "error"
                            );
                        }
                    })
                    .catch((error) => {
                        showStatus(
                            "Error generating address: " + error,
                            "error"
                        );
                    });
            }
        </script>
    </body>
</html>`