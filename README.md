# 🗺️ Jāņa Maršruta Plānotājs

SIA "Strādīgie" pārstāvis Jānis ir Rīgas Rātslaukumā un viņam ir **tikai 24 stundas**, lai apmeklētu visas 41 Latvijas pašvaldību un atgrieztos atpakaļ. Šī programma atrod īsāko iespējamo maršrutu.

---

## 📋 Uzdevuma Apraksts

Jānim ir jāapmeklē šādas pašvaldības (alfabētiskā secībā):

|                        |                          |                        |
| ---------------------- | ------------------------ | ---------------------- |
| Aizkraukles novads     | Alūksnes novads          | Augšdaugavas novads    |
| Ādažu novads           | Balvu novads             | Bauskas novads         |
| Cēsu novads            | Daugavpils valstspilsēta | Dienvidkurzemes novads |
| Dobeles novads         | Gulbenes novads          | Jelgavas novads        |
| Jelgavas valstspilsēta | Jēkabpils novads         | Jūrmalas valstspilsēta |
| Krāslavas novads       | Kuldīgas novads          | Ķekavas novads         |
| Liepājas valstspilsēta | Limbažu novads           | Līvānu novads          |
| Ludzas novads          | Madonas novads           | Mārupes novads         |
| Ogres novads           | Olaines novads           | Preiļu novads          |
| Rēzeknes novads        | Rēzeknes valstspilsēta   | Ropažu novads          |
| Salaspils novads       | Saldus novads            | Saulkrastu novads      |
| Siguldas novads        | Smiltenes novads         | Talsu novads           |
| Tukuma novads          | Valkas novads            | Valmieras novads       |
| Ventspils novads       | Ventspils valstspilsēta  |                        |

**Kopā:** 41 pašvaldība + atgriešanās Rīgā = **42 punkti maršrutā**

---

## ⚙️ Algoritms — Kā Tiek Atrasts Īsākais Maršruts

Šī ir klasiska **Komivojažiera problēma (Travelling Salesman Problem, TSP)**. Pilnīgais pārlase ir neiespējama — 41 pilsētai tas būtu **40! ≈ 10⁴⁷ varianti**.

Tāpēc risinājums izmanto trīs pakāpjus:

### 1. Haversine Formula — Attāluma Aprēķins

Attālums starp diviem ģeogrāfiskiem punktiem tiek aprēķināts pēc Haversine formulas, kas ņem vērā Zemes sfērisko formu:

```
a = sin²(ΔLat/2) + cos(lat₁) · cos(lat₂) · sin²(ΔLon/2)
d = 2R · arcsin(√a)       kur R = 6371 km
```

Rezultāts ir **taisnā līnijā mērīts attālums** pa Zemes virsmu.

### 2. Nearest Neighbor (Tuvākais Kaimiņš) — Sākotnējais Maršruts

Vienkāršs mantkārīgs algoritms sākotnējā maršruta izveidei:

```
1. Sāc no Rīgas
2. Dodies uz tuvāko vēl neapmeklēto pašvaldību
3. Atkārto, līdz visas apmeklētas
4. Atgriezies Rīgā
```

Ātrums: **O(n²)** — ātri, bet var dot suboptimālu rezultātu ar "lēcieniem".

### 3. 2-opt — Maršruta Uzlabošana

Nearest Neighbor rezultāts tiek uzlabots, novēršot šķērsojošos ceļu segmentus:

```
Kamēr ir uzlabojumi:
  Katriem diviem maršruta posmiem (i→i+1) un (j→j+1):
    Ja apmaiņa saīsina kopējo garumu:
      Apgriezt segmentu starp i+1 un j
```

Vizuāli — ja divi ceļi "krustojās kartē", pēc 2-opt tie vairs nekrustojas.

### 4. Multi-Start — Labākā Rezultāta Izvēle

Nearest Neighbor rezultāts ir atkarīgs no sākuma punkta. Tāpēc algoritms tiek palaists no **katras no 42 pilsētām** kā iekšējā sākuma punkta, un tiek izvēlēts īsākais maršruts. **Galīgais maršruts vienmēr sākas un beidzas Rīgā** — multi-start tikai uzlabo iekšējo punktu secību.

```
Labākais := ∞
Katrai pilsētai kā sākuma punktam:
    maršruts := NearestNeighbor(sākums)
    maršruts := normalizē (Rīga pirmajā vietā)
    maršruts := 2-opt(maršruts)
    Ja garums(maršruts) < Labākais:
        Labākais := maršruts
```

---

## 📊 Rezultāti

### Maršruta Secība

```
🏁 Rīga → Ķekava → Salaspils → Ogre → Ropaži → Ādaži →
Saulkrasti → Limbaži → Sigulda → Cēsis → Valmiera →
Smiltene → Valka → Alūksne → Balvi → Gulbene →
Madona → Rēzekne (valstspilsēta) → Rēzekne (novads) →
Ludza → Krāslava → Daugavpils → Augšdaugava →
Preiļi → Līvāni → Jēkabpils → Aizkraukle → Bauska →
Jelgava (valstspilsēta) → Jelgava (novads) → Dobele →
Saldus → Dienvidkurzeme → Liepāja → Kuldīga →
Ventspils (novads) → Ventspils (valstspilsēta) →
Talsi → Tukums → Jūrmala → Olaine → Mārupe → 🏁 Rīga
```

```
KOPĒJAIS MARŠRUTA GARUMS: 1506.1 km
```

Maršruts ģeogrāfiski apiet Latviju **pulksteņrādītāja virzienā**: vispirms uz austrumiem, tad caur dienvidiem, rietumiem un atpakaļ uz ziemeļiem.

---

## 🔧 Tehnoloģijas un Pieņēmumi

### Valoda

**Go (Golang)** — izmantota skaidras un lasāmas algoritma struktūras dēļ, ar statisku tipizāciju un ātru izpildi.

---

## 🚀 Palaišana

### Prasības

- Go 1.18 vai jaunāka versija

### Komandas

```bash
# Palaist tieši
go run tsp.go

# Vai kompilēt un palaist
go build -o tsp_latvia tsp.go
./tsp_latvia
```

---

## 📁 Failu Struktūra

```
.
└── tsp.go     # Galvenais fails ar visiem algoritmiem
```

Visa loģika atrodas vienā failā ar skaidri nodalītām funkcijām:

```
municipalities  →  Dati (koordinātas)
haversine()     →  Attāluma aprēķins
buildDistMatrix()  →  Attālumu matrica
nearestNeighbor()  →  Mantkārīgais sākotnējais maršruts
twoOpt()        →  Maršruta optimizācija
normalizeToRiga()  →  Rīga vienmēr pirmajā vietā
solve()         →  Multi-start orķestrācija
main()          →  Izvads
```

---
