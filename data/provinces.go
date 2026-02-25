package data

// Province represents a Thai province with geographic info.
type Province struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	NameEn string  `json:"name_en"`
	Region string  `json:"region"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}

// Provinces is the full list of 77 Thai provinces.
var Provinces = []Province{
	// Central
	{10, "กรุงเทพมหานคร", "Bangkok", "Central", 13.7563, 100.5018},
	{11, "สมุทรปราการ", "Samut Prakan", "Central", 13.5990, 100.5998},
	{12, "นนทบุรี", "Nonthaburi", "Central", 13.8622, 100.5144},
	{13, "ปทุมธานี", "Pathum Thani", "Central", 14.0208, 100.5250},
	{14, "พระนครศรีอยุธยา", "Phra Nakhon Si Ayutthaya", "Central", 14.3532, 100.5676},
	{15, "อ่างทอง", "Ang Thong", "Central", 14.5896, 100.4554},
	{16, "ลพบุรี", "Lop Buri", "Central", 14.7995, 100.6534},
	{17, "สิงห์บุรี", "Sing Buri", "Central", 14.8900, 100.3975},
	{18, "ชัยนาท", "Chai Nat", "Central", 15.1851, 100.1246},
	{19, "สระบุรี", "Saraburi", "Central", 14.5289, 100.9097},
	{60, "นครสวรรค์", "Nakhon Sawan", "Central", 15.6980, 100.0996},
	{61, "อุทัยธานี", "Uthai Thani", "Central", 15.3835, 100.0247},
	{62, "กำแพงเพชร", "Kamphaeng Phet", "Central", 16.4827, 99.5221},
	{66, "พิจิตร", "Phichit", "Central", 16.4398, 100.3488},
	{72, "สุพรรณบุรี", "Suphanburi", "Central", 14.4744, 100.1177},
	{73, "นครปฐม", "Nakhon Pathom", "Central", 13.8196, 100.0444},
	{74, "สมุทรสาคร", "Samut Sakhon", "Central", 13.5475, 100.2749},
	{75, "สมุทรสงคราม", "Samut Songkhram", "Central", 13.4098, 100.0023},
	// East
	{20, "ชลบุรี", "Chon Buri", "East", 13.3611, 100.9847},
	{21, "ระยอง", "Rayong", "East", 12.6814, 101.2816},
	{22, "จันทบุรี", "Chanthaburi", "East", 12.6112, 102.1035},
	{23, "ตราด", "Trat", "East", 12.2428, 102.5178},
	{24, "ฉะเชิงเทรา", "Chachoengsao", "East", 13.6904, 101.0779},
	{25, "ปราจีนบุรี", "Prachin Buri", "East", 14.0503, 101.3676},
	{26, "นครนายก", "Nakhon Nayok", "East", 14.2036, 101.2127},
	{27, "สระแก้ว", "Sa Kaeo", "East", 13.8237, 102.0641},
	// Northeast (Isan)
	{30, "นครราชสีมา", "Nakhon Ratchasima", "Northeast", 14.9799, 102.0977},
	{31, "บุรีรัมย์", "Buri Ram", "Northeast", 14.9952, 103.1019},
	{32, "สุรินทร์", "Surin", "Northeast", 14.8823, 103.4940},
	{33, "ศรีสะเกษ", "Si Sa Ket", "Northeast", 15.1186, 104.3219},
	{34, "อุบลราชธานี", "Ubon Ratchathani", "Northeast", 15.2448, 104.8473},
	{35, "ยโสธร", "Yasothon", "Northeast", 15.7924, 104.1451},
	{36, "ชัยภูมิ", "Chaiyaphum", "Northeast", 15.8068, 102.0318},
	{37, "อำนาจเจริญ", "Amnat Charoen", "Northeast", 15.8620, 104.6265},
	{38, "บึงกาฬ", "Bueng Kan", "Northeast", 18.3609, 103.6467},
	{39, "หนองบัวลำภู", "Nong Bua Lam Phu", "Northeast", 17.2041, 102.4417},
	{40, "ขอนแก่น", "Khon Kaen", "Northeast", 16.4419, 102.8360},
	{41, "อุดรธานี", "Udon Thani", "Northeast", 17.4138, 102.7870},
	{42, "เลย", "Loei", "Northeast", 17.4860, 101.7223},
	{43, "หนองคาย", "Nong Khai", "Northeast", 17.8783, 102.7418},
	{44, "มหาสารคาม", "Maha Sarakham", "Northeast", 16.1845, 103.3000},
	{45, "ร้อยเอ็ด", "Roi Et", "Northeast", 16.0538, 103.6520},
	{46, "กาฬสินธุ์", "Kalasin", "Northeast", 16.4314, 103.5062},
	{47, "สกลนคร", "Sakon Nakhon", "Northeast", 17.1545, 104.1348},
	{48, "นครพนม", "Nakhon Phanom", "Northeast", 17.3920, 104.7794},
	{49, "มุกดาหาร", "Mukdahan", "Northeast", 16.5434, 104.7237},
	// North
	{50, "เชียงใหม่", "Chiang Mai", "North", 18.7884, 98.9853},
	{51, "ลำพูน", "Lamphun", "North", 18.5745, 99.0087},
	{52, "ลำปาง", "Lampang", "North", 18.2888, 99.4926},
	{53, "อุตรดิตถ์", "Uttaradit", "North", 17.6200, 100.0993},
	{54, "แพร่", "Phrae", "North", 18.1446, 100.1403},
	{55, "น่าน", "Nan", "North", 18.7847, 100.7740},
	{56, "พะเยา", "Phayao", "North", 19.1664, 100.2004},
	{57, "เชียงราย", "Chiang Rai", "North", 19.9105, 99.8406},
	{58, "แม่ฮ่องสอน", "Mae Hong Son", "North", 19.2987, 97.9637},
	{63, "ตาก", "Tak", "North", 16.8798, 99.1258},
	{64, "สุโขทัย", "Sukhothai", "North", 17.0066, 99.8265},
	{65, "พิษณุโลก", "Phitsanulok", "North", 16.8211, 100.2659},
	{67, "เพชรบูรณ์", "Phetchabun", "North", 16.4189, 101.1547},
	// West
	{70, "ราชบุรี", "Ratchaburi", "West", 13.5282, 99.8134},
	{71, "กาญจนบุรี", "Kanchanaburi", "West", 14.0023, 99.5328},
	{76, "เพชรบุรี", "Phetchaburi", "West", 13.1119, 99.9488},
	{77, "ประจวบคีรีขันธ์", "Prachuap Khiri Khan", "West", 11.8126, 99.7975},
	// South
	{80, "นครศรีธรรมราช", "Nakhon Si Thammarat", "South", 8.4304, 100.0036},
	{81, "กระบี่", "Krabi", "South", 8.0863, 98.9063},
	{82, "พังงา", "Phang Nga", "South", 8.4507, 98.5264},
	{83, "ภูเก็ต", "Phuket", "South", 7.9519, 98.3381},
	{84, "สุราษฎร์ธานี", "Surat Thani", "South", 9.1382, 99.3214},
	{85, "ระนอง", "Ranong", "South", 9.9528, 98.6085},
	{86, "ชุมพร", "Chumphon", "South", 10.4930, 99.1800},
	{90, "สงขลา", "Songkhla", "South", 7.1895, 100.5950},
	{91, "สตูล", "Satun", "South", 6.6238, 100.0672},
	{92, "ตรัง", "Trang", "South", 7.5591, 99.6130},
	{93, "พัทลุง", "Phatthalung", "South", 7.6168, 100.0740},
	{94, "ปัตตานี", "Pattani", "South", 6.8653, 101.2509},
	{95, "ยะลา", "Yala", "South", 6.5413, 101.2808},
	{96, "นราธิวาส", "Narathiwat", "South", 6.4254, 101.8253},
}

// ProvinceByID returns a province by its ID, and a bool indicating found/not-found.
func ProvinceByID(id int) (Province, bool) {
	for _, p := range Provinces {
		if p.ID == id {
			return p, true
		}
	}
	return Province{}, false
}
