### 📄 File: `README.md` | 🏷️ Markdown

```markdown
# 🆔 Sentiric Identity Service

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![Language](https://img.shields.io/badge/language-Go-blue.svg)]()
[![Security](https://img.shields.io/badge/security-JWT_&_mTLS-red.svg)]()

**Sentiric Identity Service**, platformun kritik Kimlik Doğrulama (AuthN) ve Yetkilendirme (AuthZ) işlemlerini yöneten merkezi bir Control Plane bileşenidir. Tüm oturum açma, JWT üretimi ve token doğrulama mantığı burada yer alır.

## 🎯 Temel Sorumluluklar

*   **JWT Üretimi:** Geçerli kimlik bilgileri karşılığında Tenant ID ve User ID içeren güvenli JWT'ler üretir.
*   **Token Doğrulama:** Gelen JWT'lerin geçerliliğini ve süresini kontrol eder.
*   **Password Hashing:** Şifre güvenliğini yönetir (Argon2 veya Bcrypt gibi algoritmalar kullanılmalıdır).
*   **Ayrışma:** Kullanıcı verisini (`user-service`) tutmaktan ayrılmıştır.

## 🛠️ Teknoloji Yığını

*   **Dil:** Go
*   **Servisler Arası İletişim:** gRPC (Tonic)
*   **Token Standardı:** JWT (golang-jwt/jwt/v5)
*   **Bağımlılıklar:** `sentiric-contracts` v1.9.0

## 🔌 API Etkileşimleri

*   **Gelen (Sunucu):**
    *   `sentiric-api-gateway-service` (gRPC): `Authenticate`, `AuthorizeToken`.
*   **Giden (İstemci):**
    *   `sentiric-user-service` (gRPC): Şifre hash'lerini doğrulamak için (Authentication sırasında).

---
## 🏛️ Anayasal Konum

Bu servis, [Sentiric Anayasası'nın](https://github.com/sentiric/sentiric-governance) **Control Plane Layer**'ında yer alan kritik bir bileşendir.