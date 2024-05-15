## Arsitektur
<hr/>
Menggunakan Heksagonal arsitektur, dimana service ini dipisah menjadi 3 bagan besar yaitu blok kiri(interface), blok tengah(core), dan blok kanan (infrastructure).

#### Perhatikan gambar berikut:
<img src="hexa_arch.jpeg">

### Blok kiri (Interface)
Merupakan bagian antarmuka yang digunakan mengambil input data dari user(client) kemudian mengemas nya kembali untuk dikirimkan ke usecase(alur bisnis), dan mengembalikan kepada client requestor.

### Blok tengah (Core)
Merupakan bagian inti dari sebuah service(bussines logic), inti/core ini tidak asal di inject/dipanggil, tetapi wajib melalui sebuah port. Port adalah sebuah interface yang sudah di inisiasi bersamaan ketika core dibuat.

### Blok Kanan (Infrastructure)
Bagian Kanan ini biasanya menyimpan kumpulan teknologi yang digunakan dalam membangun service(bukan bussiness logic), seperti repository, third-party, dll

### Ihtisar
Oleh karena itu, berikut ini adalah susunan folder yang dibangun berdasarkan arsitektur heksagonal

app/                                # tempat inisiasi aplikasi
    app_user.go                     # tempat untuk init framework echo dan init DB
core/
    entities/                       # berisi kumpulan entities dan domain logicnya
        user.go
    port/
        user/   
            repository.go           # berisi method interface dari repository
            service.go              # berisi method interface dari service
        service/
            user.go                 # berisi service yang exec domain logic
infrastructure/
    repository/
        mysql/
            adapter/
                user_repository_adapter.go # berisi mengubah entity ke model
            models/
                user_model.go       # model yang merepresentasikan structure table
        connect.go                  # untuk init connect dengan DB
        user_repository.go          # berisi query yang di exec
    user_gateway.go                 # berisi inject koneksi DB dengan repository
interface/
    rest/
        user/  
            handler.go              # berisi handler untuk diteruskan ke service
            request.go              # payload request
            response.go             # payload response
        router.go                   # berisi route dari rest
shared/                             # berisi helper library
test/                               # berisi mock
    



