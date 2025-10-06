# 🆔 Sentiric Identity Service - Görev Listesi

Bu servisin mevcut ve gelecekteki tüm geliştirme görevleri, platformun merkezi görev yönetimi reposu olan **`sentiric-tasks`**'ta yönetilmektedir.

➡️ **[Aktif Görev Panosuna Git](https://github.com/sentiric/sentiric-tasks/blob/main/TASKS.md)**

---
Bu belge, servise özel, çok küçük ve acil görevler için geçici bir not defteri olarak kullanılabilir.

## Faz 1: Minimal İşlevsellik (INFRA-02)
- [x] Temel Go projesi, JWT kütüphaneleri ve Dockerfile oluşturuldu.
- [x] gRPC sunucusu iskeleti (`Authenticate`, `AuthorizeToken`) eklendi.
- [x] Temel JWT oluşturma ve doğrulama mantığı (`internal/token`) eklendi.
- [ ] User Service'e şifre doğrulama için gRPC istemcisi entegre edilecek. (CONTROL-ID-01)