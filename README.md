# get_google_trends_api

This is just a simple go program to get the Google Trends API data, and then do like take a part 
the XML as the returned data. And be served like a simple list such the name of news, the link of news and etc.

This is how it looks like:

Kita akan mengambil data API dari Google Trends 
---------------------------------------------------
# 1
Search Term: Leverkusen
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Leverkusen
Headline: Prediksi Atalanta Vs Leverkusen: Die Werkself Bakal Jalani Laga ...
Link to article: https://sport.detik.com/sepakbola/uefa/d-7353170/prediksi-atalanta-vs-leverkusen-die-werkself-bakal-jalani-laga-berat
--------------------------
# 2
Search Term: 23 Mei 2024 hari apa
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#23%20Mei%202024%20hari%20apa
Headline: Besok 23 Mei 2024 Hari Apa? Ini Peringatan dan Jadwal Liburnya
Link to article: https://www.detik.com/sulsel/berita/d-7353690/besok-23-mei-2024-hari-apa-ini-peringatan-dan-jadwal-liburnya
--------------------------
# 3
Search Term: Ryukyu FC vs Cerezo Osaka
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Ryukyu%20FC%20vs%20Cerezo%20Osaka
Headline: Duel Justin Hubner Berbuah Kartu Merah, Cerezo Osaka ...
Link to article: https://www.bolasport.com/read/314092815/duel-justin-hubner-berbuah-kartu-merah-cerezo-osaka-tumbangkan-fc-ryukyu
--------------------------
# 4
Search Term: Pegi
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Pegi
Headline: Begini Tampang Pegi, DPO Kasus Pembunuhan Vina Cirebon
Link to article: https://news.detik.com/berita/d-7353818/begini-tampang-pegi-dpo-kasus-pembunuhan-vina-cirebon
--------------------------
# 5
Search Term: Rekrutmen BUMN
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Rekrutmen%20BUMN
Headline: Hasil Tes Online 1 Rekrutmen BUMN Diumumkan Hari Ini, Begini ...
Link to article: https://www.kompas.com/tren/read/2024/05/22/071500565/hasil-tes-online-1-rekrutmen-bumn-diumumkan-hari-ini-begini-cara-ceknya
--------------------------
# 6
Search Term: BUMN
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#BUMN
Headline: Stafsus Menteri BUMN sebut akar masalah Indofarma dari anak ...
Link to article: https://www.antaranews.com/berita/4116864/stafsus-menteri-bumn-sebut-akar-masalah-indofarma-dari-anak-usahanya
--------------------------
# 7
Search Term: Furiosa
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Furiosa
Headline: Jadwal Bioskop Bali Rabu 22 Mei 2024, Ada Film Furiosa: A Mad ...
Link to article: https://www.detik.com/bali/berita/d-7351145/jadwal-bioskop-bali-rabu-22-mei-2024-ada-film-furiosa-a-mad-max-saga
--------------------------
# 8
Search Term: Eminem
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Eminem
Headline: Hadir di Pernikahan Putrinya, Gaya Eminem Berkacamata Hitam ...
Link to article: https://wolipop.detik.com/entertainment-news/d-7353398/hadir-di-pernikahan-putrinya-gaya-eminem-berkacamata-hitam-jadi-sorotan
--------------------------
# 9
Search Term: Tuhan Izinkan Aku Berdosa
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Tuhan%20Izinkan%20Aku%20Berdosa
Headline: Tuhan Izinkan Aku Berdosa Sempat Ditolak Banyak Produser dan OTT
Link to article: https://www.detik.com/pop/movie/d-7352745/tuhan-izinkan-aku-berdosa-sempat-ditolak-banyak-produser-dan-ott
--------------------------
# 10
Search Term: Pochettino
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Pochettino
Headline: Aneh Rasanya Lihat Chelsea dan Pochettino Pisah
Link to article: https://sport.detik.com/sepakbola/liga-inggris/d-7353769/aneh-rasanya-lihat-chelsea-dan-pochettino-pisah
--------------------------
# 11
Search Term: Saldo
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Saldo
Headline: Cara Cek Saldo BSI lewat ATM dan Mobile Banking
Link to article: https://money.kompas.com/read/2024/05/22/235251326/cara-cek-saldo-bsi-lewat-atm-dan-mobile-banking
--------------------------
# 12
Search Term: NAC Breda
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#NAC%20Breda
Headline: Hasil Playoff Promosi Liga Belanda: Pesaing Tim Rafael Struick ...
Link to article: https://kalteng.tribunnews.com/2024/05/22/hasil-playoff-promosi-liga-belanda-pesaing-tim-rafael-struick-nac-breda-vs-fc-emmen-hanya-imbang
--------------------------
# 13
Search Term: UNS
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#UNS
Headline: Mahasiswa MBKM UNS dan Kelompok Tani di Desa Kemudo ...
Link to article: https://uns.ac.id/id/uns-update/mahasiswa-mbkm-uns-dan-kelompok-tani-di-desa-kemudo-klaten-ciptakan-solusi-pengendalian-hama-dari-buah-maja.html
--------------------------
# 14
Search Term: Boss Level
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Boss%20Level
Headline: Sinopsis Film Boss Level, Mengungkap Misteri Kematian Mantan ...
Link to article: https://www.detik.com/jateng/budaya/d-7353586/sinopsis-film-boss-level-mengungkap-misteri-kematian-mantan-pasukan-khusus
--------------------------
# 15
Search Term: 23 Mei 2024 libur Apa
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#23%20Mei%202024%20libur%20Apa
Headline: Prakiraan Cuaca Libur Hari Raya Waisak 23 Mei 2024 di 34 Kota ...
Link to article: https://news.detik.com/berita/d-7353574/prakiraan-cuaca-libur-hari-raya-waisak-23-mei-2024-di-34-kota-besar
--------------------------
# 16
Search Term: Gempa hari ini
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Gempa%20hari%20ini
Headline: Rentetan Gempa Hari Ini Guncang Beberapa Wilayah, Terkuat Di ...
Link to article: https://rm.id/baca-berita/nasional/221486/rentetan-gempa-hari-ini-guncang-beberapa-wilayah-terkuat-di-sulut-m-56
--------------------------
# 17
Search Term: Serie B
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Serie%20B
Headline: Venezia OTW Final Playoff Serie B, Jay Idzes Bisa Absen Perkuat ...
Link to article: https://www.bola.com/indonesia/read/5600798/venezia-otw-final-playoff-serie-b-jay-idzes-bisa-absen-perkuat-timnas-indonesia-vs-tanzania
--------------------------
# 18
Search Term: Toni Kroos
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Toni%20Kroos
Headline: Toni Kroos Pensiun dari Lapangan Hijau, Banting Setir Jadi Bos ...
Link to article: https://www.detik.com/properti/berita/d-7352987/toni-kroos-pensiun-dari-lapangan-hijau-banting-setir-jadi-bos-properti
--------------------------
# 19
Search Term: Wuthering Waves
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Wuthering%20Waves
Headline: Download Wuthering Waves: Game Open World RPG Terbaru ...
Link to article: https://padek.jawapos.com/gaya-hidup/2364677833/download-wuthering-waves-game-open-world-rpg-terbaru-cerita-pasca-apokaliptik-yang-menarik
--------------------------
# 20
Search Term: Singapore Airlines
Link to trend: https://trends.google.com/trends/trendingsearches/daily?geo=ID#Singapore%20Airlines
Headline: Ungkapan Belasungkawa CEO Singapore Airlines Usai SQ321 ...
Link to article: https://news.detik.com/berita/d-7352588/ungkapan-belasungkawa-ceo-singapore-airlines-usai-sq321-alami-turbulensi
--------------------------