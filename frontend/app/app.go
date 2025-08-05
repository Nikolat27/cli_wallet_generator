package app

import (
	"encoding/json"
	"fmt"
	"go_wallet_generator/address"
	"go_wallet_generator/crypto"
	"go_wallet_generator/wallet"
	"net/http"
	"strings"
)

type WalletApp struct {
	// Web server components
}

type WalletResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type WalletRequest struct {
	Name string `json:"name"`
	Coin string `json:"coin"`
}

func NewWalletApp() *WalletApp {
	return &WalletApp{}
}

func (wa *WalletApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.URL.Path {
	case "/":
		wa.serveMainPage(w, r)
	case "/api/wallets":
		wa.handleWalletsAPI(w, r)
	case "/api/addresses":
		wa.handleAddressesAPI(w, r)
	case "/api/create-wallet":
		wa.handleCreateWallet(w, r)
	case "/api/delete-wallet":
		wa.handleDeleteWallet(w, r)
	case "/api/generate-address":
		wa.handleGenerateAddress(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (wa *WalletApp) serveMainPage(w http.ResponseWriter, r *http.Request) {
	// Serve the embedded HTML content
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(EmbeddedHTML))
}

func (wa *WalletApp) handleWalletsAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wallets, err := wa.getWallets()
	if err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(WalletResponse{
		Success: true,
		Data:    wallets,
	})
}

func (wa *WalletApp) handleAddressesAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	walletName := r.URL.Query().Get("wallet")
	if walletName == "" {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Wallet name is required",
		})
		return
	}

	addresses, err := wa.getAddresses(walletName)
	if err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(WalletResponse{
		Success: true,
		Data:    addresses,
	})
}

func (wa *WalletApp) handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req WalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Wallet name is required",
		})
		return
	}

	walletInstance := wallet.Constructor()
	walletInstance.Name = req.Name

	if err := walletInstance.CreateWallet(); err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Get the wallet instance to retrieve the mnemonic
	walletData, err := walletInstance.GetWalletInstance()
	if err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Error retrieving wallet data: " + err.Error(),
		})
		return
	}

	// Decrypt the mnemonic to show to user
	mnemonic, err := wa.decryptMnemonic(walletData.Mnemonic)
	if err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Error decrypting mnemonic: " + err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(WalletResponse{
		Success: true,
		Message: "Wallet created successfully! Mnemonic copied to clipboard.",
		Data: map[string]interface{}{
			"mnemonic":    string(mnemonic),
			"wallet_name": req.Name,
		},
	})
}

func (wa *WalletApp) handleDeleteWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req WalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Wallet name is required",
		})
		return
	}

	walletInstance := wallet.Constructor()
	walletInstance.Name = req.Name

	if err := walletInstance.DeleteWallet(); err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(WalletResponse{
		Success: true,
		Message: "Wallet deleted successfully",
	})
}

func (wa *WalletApp) handleGenerateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Wallet string `json:"wallet"`
		Coin   string `json:"coin"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Wallet == "" {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Wallet name is required",
		})
		return
	}

	if req.Coin == "" {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: "Coin type is required",
		})
		return
	}

	_, err := address.GenerateAndStoreAddress(req.Wallet, req.Coin)
	if err != nil {
		json.NewEncoder(w).Encode(WalletResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(WalletResponse{
		Success: true,
		Message: fmt.Sprintf("%s address generated successfully", strings.ToUpper(req.Coin)),
	})
}

func (wa *WalletApp) getWallets() ([]map[string]interface{}, error) {
	w := wallet.Constructor()
	wallets, err := w.ListWallets()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, w := range wallets {
		result = append(result, map[string]interface{}{
			"name":       w.Name,
			"created_at": w.CreatedAt.Format("2006-01-02"),
		})
	}

	// Always return an array, even if empty
	return result, nil
}

func (wa *WalletApp) getAddresses(walletName string) ([]map[string]interface{}, error) {
	addresses, err := address.RetrieveAddressList(walletName)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, addr := range addresses {
		result = append(result, map[string]interface{}{
			"coin":    addr.Coin,
			"address": addr.Address,
		})
	}

	return result, nil
}

func (wa *WalletApp) decryptMnemonic(encryptedMnemonic string) ([]byte, error) {
	return crypto.DecryptBase64(encryptedMnemonic)
}
