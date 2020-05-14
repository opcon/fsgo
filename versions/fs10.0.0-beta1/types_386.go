// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs fs10.0.0-beta1/types.go

package fs

type (
	Bank_set_mon struct {
		ActiveBank struct {
			ActiveBank [33]byte `json:"active_bank"`
			padCgo0    [3]byte
			State      M5state `json:"state"`
		} `json:"active_bank"`
		ActiveVsn struct {
			ActiveVsn [33]byte `json:"active_vsn"`
			padCgo0   [3]byte
			State     M5state `json:"state"`
		} `json:"active_vsn"`
		InactiveBank struct {
			InactiveBank [33]byte `json:"inactive_bank"`
			padCgo0      [3]byte
			State        M5state `json:"state"`
		} `json:"inactive_bank"`
		InactiveVsn struct {
			InactiveVsn [33]byte `json:"inactive_vsn"`
			padCgo0     [3]byte
			State       M5state `json:"state"`
		} `json:"inactive_vsn"`
	}
	Bbc_cmd struct {
		Freq   int32    `json:"freq"`
		Source int32    `json:"source"`
		Bw     [2]int32 `json:"bw"`
		Bwcomp [2]int32 `json:"bwcomp"`
		Gain   struct {
			Mode  int32    `json:"mode"`
			Value [2]int32 `json:"value"`
			Old   int32    `json:"old"`
		} `json:"gain"`
		Avper int32 `json:"avper"`
	}
	Bbc_mon struct {
		Lock   int32     `json:"lock"`
		Pwr    [2]uint32 `json:"pwr"`
		Serno  int32     `json:"serno"`
		Timing int32     `json:"timing"`
	}
	Bs struct {
		ImageRejectFilter uint32 `json:"image_reject_filter"`
		Level             Servo  `json:"level"`
		Offset            Servo  `json:"offset"`
		MagnStats         uint32 `json:"magn_stats"`
		Flip64MHzOut      uint32 `json:"flip_64MHz_out"`
		DigitalFormat     uint32 `json:"digital_format"`
		FlipInput         uint32 `json:"flip_input"`
		PHilbertNo        byte   `json:"p_hilbert_no"`
		NHilbertNo        byte   `json:"n_hilbert_no"`
		padCgo0           [2]byte
		SubBand           uint32 `json:"sub_band"`
		QFirNo            byte   `json:"q_fir_no"`
		IFirNo            byte   `json:"i_fir_no"`
		ClockDecimation   byte   `json:"clock_decimation"`
		padCgo1           [1]byte
		AddSub            Mux  `json:"add_sub"`
		UsbMux            Mux  `json:"usb_mux"`
		LsbMux            Mux  `json:"lsb_mux"`
		UsbThreshold      byte `json:"usb_threshold"`
		LsbThreshold      byte `json:"lsb_threshold"`
		padCgo2           [2]byte
		UsbServo          Servo  `json:"usb_servo"`
		LsbServo          Servo  `json:"lsb_servo"`
		FlipUsb           uint32 `json:"flip_usb"`
		FlipLsb           uint32 `json:"flip_lsb"`
		Monitor           Mux    `json:"monitor"`
		Digout            Digout `json:"digout"`
	}
	Calrx_cmd struct {
		File    [65]byte `json:"file"`
		padCgo0 [3]byte
		Type    int32      `json:"type"`
		Lo      [2]float64 `json:"lo"`
	}
	Capture_mon struct {
		Qa struct {
			Drive int32 `json:"drive"`
			Chan  int32 `json:"chan"`
		} `json:"qa"`
		General struct {
			Word1 uint32 `json:"word1"`
			Word2 uint32 `json:"word2"`
		} `json:"general"`
		Time struct {
			Word3 uint32 `json:"word3"`
			Word4 uint32 `json:"word4"`
		} `json:"time"`
	}
	Clock_set_cmd struct {
		Freq struct {
			Freq  int32   `json:"freq"`
			State M5state `json:"state"`
		} `json:"freq"`
		Source struct {
			Source  [33]byte `json:"source"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"source"`
		ClockGen struct {
			ClockGen float64 `json:"clock_gen"`
			State    M5state `json:"state"`
		} `json:"clock_gen"`
	}
	Cmd_ds struct {
		Name    *byte `json:"name"`
		Equal   byte  `json:"equal"`
		padCgo0 [3]byte
		Argv    [100]*byte `json:"argv"`
	}
	Das struct {
		DsMnem         [3]byte `json:"ds_mnem"`
		padCgo0        [1]byte
		Ifp            [2]Ifp  `json:"ifp"`
		VoltageP5VIfp1 float32 `json:"voltage_p5V_ifp1"`
		VoltageP5VIfp2 float32 `json:"voltage_p5V_ifp2"`
		VoltageM5d2V   float32 `json:"voltage_m5d2V"`
		VoltageP9V     float32 `json:"voltage_p9V"`
		VoltageM9V     float32 `json:"voltage_m9V"`
		VoltageP15V    float32 `json:"voltage_p15V"`
		VoltageM15V    float32 `json:"voltage_m15V"`
	}
	Data_check_mon struct {
		Missing struct {
			Missing int64   `json:"missing"`
			State   M5state `json:"state"`
		} `json:"missing"`
		Mode struct {
			Mode    [33]byte `json:"mode"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"mode"`
		Submode struct {
			Submode [33]byte `json:"submode"`
			padCgo0 [3]byte
			First   int32   `json:"first"`
			State   M5state `json:"state"`
		} `json:"submode"`
		Time struct {
			Time  M5time  `json:"time"`
			Bad   int32   `json:"bad"`
			State M5state `json:"state"`
		} `json:"time"`
		Offset struct {
			Offset int32   `json:"offset"`
			Size   int32   `json:"size"`
			State  M5state `json:"state"`
		} `json:"offset"`
		Period struct {
			Period M5time  `json:"period"`
			State  M5state `json:"state"`
		} `json:"period"`
		Bytes struct {
			Bytes int32   `json:"bytes"`
			State M5state `json:"state"`
		} `json:"bytes"`
		Source struct {
			Source  [33]byte `json:"source"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"source"`
		Start struct {
			Start M5time  `json:"start"`
			State M5state `json:"state"`
		} `json:"start"`
		Code struct {
			Code  int32   `json:"code"`
			State M5state `json:"state"`
		} `json:"code"`
		Frames struct {
			Frames int32   `json:"frames"`
			State  M5state `json:"state"`
		} `json:"frames"`
		Header struct {
			Header M5time  `json:"header"`
			State  M5state `json:"state"`
		} `json:"header"`
		Total struct {
			Total float32 `json:"total"`
			State M5state `json:"state"`
		} `json:"total"`
		Byte struct {
			Byte  int32   `json:"byte"`
			State M5state `json:"state"`
		} `json:"byte"`
	}
	Data_valid_cmd struct {
		UserDv   int32 `json:"user_dv"`
		PbEnable int32 `json:"pb_enable"`
	}
	Dbbc3_bbcnn_cmd struct {
		Freq   uint32 `json:"freq"`
		Source int32  `json:"source"`
		Bw     int32  `json:"bw"`
		Avper  int32  `json:"avper"`
	}
	Dbbc3_bbcnn_mon struct {
		Agc   int32     `json:"agc"`
		Gain  [2]int32  `json:"gain"`
		Tpon  [2]uint32 `json:"tpon"`
		Tpoff [2]uint32 `json:"tpoff"`
	}
	Dbbc3_cont_cal_cmd struct {
		Mode     int32 `json:"mode"`
		Polarity int32 `json:"polarity"`
		Freq     int32 `json:"freq"`
		Option   int32 `json:"option"`
		Samples  int32 `json:"samples"`
	}
	Dbbc3_iftpx_mon struct {
		Tp  uint32 `json:"tp"`
		Off uint32 `json:"off"`
		On  uint32 `json:"on"`
	}
	Dbbc3_ifx_cmd struct {
		Input      int32  `json:"input"`
		Att        int32  `json:"att"`
		Agc        int32  `json:"agc"`
		TargetNull int32  `json:"target_null"`
		Target     uint32 `json:"target"`
	}
	Dbbc3_ifx_mon struct {
		Tp uint32 `json:"tp"`
	}
	Dbbc_cont_cal_cmd struct {
		Mode     int32 `json:"mode"`
		Polarity int32 `json:"polarity"`
		Freq     int32 `json:"freq"`
		Option   int32 `json:"option"`
		Samples  int32 `json:"samples"`
	}
	Dbbcform_cmd struct {
		Mode int32 `json:"mode"`
		Test int32 `json:"test"`
	}
	Dbbcgain_cmd struct {
		Bbc    int32 `json:"bbc"`
		State  int32 `json:"state"`
		GainU  int32 `json:"gainU"`
		GainL  int32 `json:"gainL"`
		Target int32 `json:"target"`
	}
	Dbbcgain_mon struct {
		State  int32 `json:"state"`
		Target int32 `json:"target"`
	}
	Dbbcifx_cmd struct {
		Input      int32  `json:"input"`
		Att        int32  `json:"att"`
		Agc        int32  `json:"agc"`
		Filter     int32  `json:"filter"`
		TargetNull int32  `json:"target_null"`
		Target     uint32 `json:"target"`
	}
	Dbbcifx_mon struct {
		Tp uint32 `json:"tp"`
	}
	Dbbcnn_cmd struct {
		Freq   uint32 `json:"freq"`
		Source int32  `json:"source"`
		Bw     int32  `json:"bw"`
		Avper  int32  `json:"avper"`
	}
	Dbbcnn_mon struct {
		Agc   int32     `json:"agc"`
		Gain  [2]int32  `json:"gain"`
		Tpon  [2]uint32 `json:"tpon"`
		Tpoff [2]uint32 `json:"tpoff"`
	}
	Dbbc_pfbx_mon struct {
		Counts   [17]int32 `json:"counts"`
		Overflow int32     `json:"overflow"`
	}
	Dbbcvsi_clk_mon struct {
		Clk struct {
			VsiClk int32   `json:"vsi_clk"`
			State  M5state `json:"state"`
		} `json:"vsi_clk"`
	}
	Dbbc_vsix_cmd struct {
		Core [16]int32 `json:"core"`
		Chan [16]int32 `json:"chan"`
	}
	Digout struct {
		Setting  uint32 `json:"setting"`
		Mode     uint32 `json:"mode"`
		Tristate uint32 `json:"tristate"`
	}
	Disk2file_cmd struct {
		ScanLabel struct {
			ScanLabel [65]byte `json:"scan_label"`
			padCgo0   [3]byte
			State     M5state `json:"state"`
		} `json:"scan_label"`
		Destination struct {
			Destination [129]byte `json:"destination"`
			padCgo0     [3]byte
			State       M5state `json:"state"`
		} `json:"destination"`
		Start struct {
			Start   [33]byte `json:"start"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"start"`
		End struct {
			End     [33]byte `json:"end"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"end"`
		Options struct {
			Options [33]byte `json:"options"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"options"`
	}
	Disk2file_mon struct {
		Option struct {
			Option  [33]byte `json:"option"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"option"`
		StartByte struct {
			StartByte int64   `json:"start_byte"`
			State     M5state `json:"state"`
		} `json:"start_byte"`
		EndByte struct {
			EndByte int64   `json:"end_byte"`
			State   M5state `json:"state"`
		} `json:"end_byte"`
		Status struct {
			Status  [33]byte `json:"status"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"status"`
		Current struct {
			Current int64   `json:"current"`
			State   M5state `json:"state"`
		} `json:"current"`
		ScanNumber struct {
			ScanNumber int32   `json:"scan_number"`
			State      M5state `json:"state"`
		} `json:"scan_number"`
	}
	Disk_pos_mon struct {
		Record struct {
			Record int64   `json:"record"`
			State  M5state `json:"state"`
		} `json:"record"`
		Play struct {
			Play  int64   `json:"play"`
			State M5state `json:"state"`
		} `json:"play"`
		Stop struct {
			Stop  int64   `json:"stop"`
			State M5state `json:"state"`
		} `json:"stop"`
	}
	Disk_record_cmd struct {
		Record struct {
			Record int32   `json:"record"`
			State  M5state `json:"state"`
		} `json:"record"`
		Label struct {
			Label   [65]byte `json:"label"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"label"`
	}
	Disk_record_mon struct {
		Status struct {
			Status  [33]byte `json:"status"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"status"`
		Scan struct {
			Scan  int32   `json:"scan"`
			State M5state `json:"state"`
		} `json:"scan"`
	}
	Disk_serial_mon struct {
		Count  int32 `json:"count"`
		Serial [16]struct {
			Serial  [33]byte `json:"serial"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"serial"`
	}
	Dist_cmd struct {
		Atten [2]int32 `json:"atten"`
		Input [2]int32 `json:"input"`
		Avper int32    `json:"avper"`
		Old   [2]int32 `json:"old"`
	}
	Dist_mon struct {
		Serial int32     `json:"serial"`
		Timing int32     `json:"timing"`
		Totpwr [2]uint32 `json:"totpwr"`
	}
	Dot_mon struct {
		Time struct {
			Time    [33]byte `json:"time"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"time"`
		Status struct {
			Status  [33]byte `json:"status"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"status"`
		FHGStatus struct {
			FHGStatus [33]byte `json:"FHG_status"`
			padCgo0   [3]byte
			State     M5state `json:"state"`
		} `json:"FHG_status"`
		OSTime struct {
			OSTime  [33]byte `json:"OS_time"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"OS_time"`
		DOTOSTimeDiff struct {
			DOTOSTimeDiff [33]byte `json:"DOT_OS_time_diff"`
			padCgo0       [3]byte
			State         M5state `json:"state"`
		} `json:"DOT_OS_time_diff"`
	}
	Dqa_chan struct {
		Bbc     int32   `json:"bbc"`
		Track   int32   `json:"track"`
		Amp     float32 `json:"amp"`
		Phase   float32 `json:"phase"`
		Parity  uint32  `json:"parity"`
		CrccA   uint32  `json:"crcc_a"`
		CrccB   uint32  `json:"crcc_b"`
		Resync  uint32  `json:"resync"`
		Nosync  uint32  `json:"nosync"`
		NumBits uint32  `json:"num_bits"`
	}
	Dqa_cmd struct {
		Dur int32 `json:"dur"`
	}
	Dqa_mon struct {
		A Dqa_chan `json:"a"`
		B Dqa_chan `json:"b"`
	}
	Ds_cmd struct {
		Type    uint16  `json:"type"`
		Mnem    [3]byte `json:"mnem"`
		padCgo0 [1]byte
		Cmd     uint16 `json:"cmd"`
		Data    uint16 `json:"data"`
	}
	Ds_mon struct {
		Resp uint16  `json:"resp"`
		Data [2]byte `json:"data"`
	}
	Fila10g_mode_cmd struct {
		Mask2 struct {
			Mask2 uint32  `json:"mask2"`
			State M5state `json:"state"`
		} `json:"mask2"`
		Mask1 struct {
			Mask1 uint32  `json:"mask1"`
			State M5state `json:"state"`
		} `json:"mask1"`
		Decimate struct {
			Decimate int32   `json:"decimate"`
			State    M5state `json:"state"`
		} `json:"decimate"`
		Disk struct {
			Disk  int32   `json:"disk"`
			State M5state `json:"state"`
		} `json:"disk"`
	}
	Fila10g_mode_mon struct {
		Clockrate struct {
			Clockrate int32   `json:"clockrate"`
			State     M5state `json:"state"`
		} `json:"clockrate"`
	}
	Flux_ds struct {
		Name    [11]byte   `json:"name"`
		Type    byte       `json:"type"`
		Fmin    float32    `json:"fmin"`
		Fmax    float32    `json:"fmax"`
		Fcoeff  [3]float32 `json:"fcoeff"`
		Size    float32    `json:"size"`
		Model   byte       `json:"model"`
		padCgo0 [3]byte
		Mcoeff  [6]float32 `json:"mcoeff"`
	}
	Form4_cmd struct {
		Mode     int32         `json:"mode"`
		Rate     int32         `json:"rate"`
		Enable   [2]uint32     `json:"enable"`
		Aux      [2]uint32     `json:"aux"`
		Codes    [64]int32     `json:"codes"`
		Bits     int32         `json:"bits"`
		Fan      int32         `json:"fan"`
		Barrel   int32         `json:"barrel"`
		Modulate int32         `json:"modulate"`
		Last     int32         `json:"last"`
		Synch    int32         `json:"synch"`
		Roll     [16][64]int32 `json:"roll"`
		StartMap int32         `json:"start_map"`
		EndMap   int32         `json:"end_map"`
		A2d      [16][64]int32 `json:"a2d"`
	}
	Form4_mon struct {
		Status  int32 `json:"status"`
		Error   int32 `json:"error"`
		RackIds int32 `json:"rack_ids"`
		Version int32 `json:"version"`
	}
	Fscom struct {
		Iclbox      int32        `json:"iclbox"`
		Iclopr      int32        `json:"iclopr"`
		Nums        [40]int32    `json:"nums"`
		AZOFF       float32      `json:"AZOFF"`
		DECOFF      float32      `json:"DECOFF"`
		ELOFF       float32      `json:"ELOFF"`
		Ibmat       int32        `json:"ibmat"`
		Ibmcb       int32        `json:"ibmcb"`
		ICAPTP      [2]int32     `json:"ICAPTP"`
		IRDYTP      [2]int32     `json:"IRDYTP"`
		IRENVC      int32        `json:"IRENVC"`
		ILOKVC      int32        `json:"ILOKVC"`
		ITRAKA      [2]int32     `json:"ITRAKA"`
		ITRAKB      [2]int32     `json:"ITRAKB"`
		TPIVC       [15]uint32   `json:"TPIVC"`
		ISTPTP      [2]float32   `json:"ISTPTP"`
		ITACTP      [2]float32   `json:"ITACTP"`
		KHALT       int32        `json:"KHALT"`
		KECHO       int32        `json:"KECHO"`
		KENASTK     [2][2]int32  `json:"KENASTK"`
		INEXT       [3]int32     `json:"INEXT"`
		RAOFF       float32      `json:"RAOFF"`
		XOFF        float32      `json:"XOFF"`
		YOFF        float32      `json:"YOFF"`
		LLOG        [8]byte      `json:"LLOG"`
		LNEWPR      [8]byte      `json:"LNEWPR"`
		LNEWSK      [8]byte      `json:"LNEWSK"`
		LPRC        [8]byte      `json:"LPRC"`
		LSTP        [8]byte      `json:"LSTP"`
		LSKD        [8]byte      `json:"LSKD"`
		LEXPER      [8]byte      `json:"LEXPER"`
		LFEETFS     [2][6]byte   `json:"LFEET_FS"`
		Lgen        [2][2]int16  `json:"lgen"`
		ICHK        [23]int32    `json:"ICHK"`
		Tempwx      float32      `json:"tempwx"`
		Humiwx      float32      `json:"humiwx"`
		Preswx      float32      `json:"preswx"`
		Speedwx     float32      `json:"speedwx"`
		Directionwx int32        `json:"directionwx"`
		Ep1950      float32      `json:"ep1950"`
		Epoch       float32      `json:"epoch"`
		Cablev      float32      `json:"cablev"`
		Height      float32      `json:"height"`
		Ra50        float64      `json:"ra50"`
		Dec50       float64      `json:"dec50"`
		Radat       float64      `json:"radat"`
		Decdat      float64      `json:"decdat"`
		Alat        float64      `json:"alat"`
		Wlong       float64      `json:"wlong"`
		Systmp      [264]float32 `json:"systmp"`
		Ldsign      int32        `json:"ldsign"`
		Lfreqv      [90]byte     `json:"lfreqv"`
		Lnaant      [8]byte      `json:"lnaant"`
		Lsorna      [10]byte     `json:"lsorna"`
		Idevant     [64]byte     `json:"idevant"`
		Idevgpib    [64]byte     `json:"idevgpib"`
		Idevlog     [64][5]byte  `json:"idevlog"`
		Ndevlog     int32        `json:"ndevlog"`
		Imodfm      int32        `json:"imodfm"`
		Ipashd      [2][2]int32  `json:"ipashd"`
		Iratfm      int32        `json:"iratfm"`
		Ispeed      [2]int32     `json:"ispeed"`
		Idirtp      [2]int32     `json:"idirtp"`
		Cips        [2]int32     `json:"cips"`
		BitDensity  [2]int32     `json:"bit_density"`
		Ienatp      [2]int32     `json:"ienatp"`
		Inp1if      int32        `json:"inp1if"`
		Inp2if      int32        `json:"inp2if"`
		Ionsor      int32        `json:"ionsor"`
		Imaxtpsd    [2]int32     `json:"imaxtpsd"`
		Iskdtpsd    [2]int32     `json:"iskdtpsd"`
		Motorv      [2]float32   `json:"motorv"`
		Inscint     [2]float32   `json:"inscint"`
		Inscsl      [2]float32   `json:"inscsl"`
		Outscint    [2]float32   `json:"outscint"`
		Outscsl     [2]float32   `json:"outscsl"`
		Itpthick    [2]int32     `json:"itpthick"`
		Wrvolt      [2]float32   `json:"wrvolt"`
		Capstan     [2]int32     `json:"capstan"`
		Go          struct {
			Allocated int32       `json:"allocated"`
			Name      [32][5]byte `json:"name"`
		} `json:"go"`
		Sem struct {
			Allocated int32       `json:"allocated"`
			Name      [32][5]byte `json:"name"`
		} `json:"sem"`
		Check struct {
			Bbc       [16]int32   `json:"bbc"`
			BbcTime   [16]int32   `json:"bbc_time"`
			Dist      [2]int32    `json:"dist"`
			Vform     int32       `json:"vform"`
			FmCnTm    int32       `json:"fm_cn_tm"`
			Rec       [2]int32    `json:"rec"`
			Vkrepro   [2]int32    `json:"vkrepro"`
			Vkenable  [2]int32    `json:"vkenable"`
			Vkmove    [2]int32    `json:"vkmove"`
			Systracks [2]int32    `json:"systracks"`
			RcMvTm    [2]int32    `json:"rc_mv_tm"`
			Vklowtape [2]int32    `json:"vklowtape"`
			Vkload    [2]int32    `json:"vkload"`
			RcLdTm    [2]int32    `json:"rc_ld_tm"`
			S2rec     S2rec_check `json:"s2rec"`
			K4rec     K4rec_check `json:"k4rec"`
			Ifp       [4]int32    `json:"ifp"`
			IfpTime   [4]int32    `json:"ifp_time"`
			DbbcForm  int32       `json:"dbbc_form"`
		} `json:"check"`
		Stcnm   [4][2]byte  `json:"stcnm"`
		Stchk   [4]int32    `json:"stchk"`
		Dist    [2]Dist_cmd `json:"dist"`
		Bbc     [16]Bbc_cmd `json:"bbc"`
		Tpi     [264]int32  `json:"tpi"`
		Tpical  [264]int32  `json:"tpical"`
		Tpizero [264]int32  `json:"tpizero"`
		Equip   struct {
			Rack        int32    `json:"rack"`
			Drive       [2]int32 `json:"drive"`
			DriveType   [2]int32 `json:"drive_type"`
			RackType    int32    `json:"rack_type"`
			WxMet       int32    `json:"wx_met"`
			WxHost      [65]byte `json:"wx_host"`
			padCgo0     [3]byte
			Mk4syncDflt int32 `json:"mk4sync_dflt"`
		} `json:"equip"`
		KlvdtFs    [2]int32          `json:"klvdt_fs"`
		Vrepro     [2]Vrepro_cmd     `json:"vrepro"`
		Vform      Vform_cmd         `json:"vform"`
		Venable    [2]Venable_cmd    `json:"venable"`
		Systracks  [2]Systracks_cmd  `json:"systracks"`
		Dqa        Dqa_cmd           `json:"dqa"`
		UserInfo   User_info_cmd     `json:"user_info"`
		S2st       S2st_cmd          `json:"s2st"`
		S2RecState int32             `json:"s2_rec_state"`
		RecMode    Rec_mode_cmd      `json:"rec_mode"`
		DataValid  [2]Data_valid_cmd `json:"data_valid"`
		S2label    S2label_cmd       `json:"s2label"`
		padCgo0    [3]byte
		Form4      Form4_cmd `json:"form4"`
		Diaman     float32   `json:"diaman"`
		Slew1      float32   `json:"slew1"`
		Slew2      float32   `json:"slew2"`
		Lolim1     float32   `json:"lolim1"`
		Lolim2     float32   `json:"lolim2"`
		Uplim1     float32   `json:"uplim1"`
		Uplim2     float32   `json:"uplim2"`
		Refreq     float32   `json:"refreq"`
		I70kch     int32     `json:"i70kch"`
		I20kch     int32     `json:"i20kch"`
		Time       struct {
			Rate      [2]float32 `json:"rate"`
			Offset    [2]int32   `json:"offset"`
			Epoch     [2]int32   `json:"epoch"`
			Span      [2]int32   `json:"span"`
			SecsOff   int32      `json:"secs_off"`
			Index     int32      `json:"index"`
			Icomputer [2]int32   `json:"icomputer"`
			Model     byte       `json:"model"`
			padCgo0   [3]byte
			TicksOff  uint32 `json:"ticks_off"`
			UsecsOff  int32  `json:"usecs_off"`
			InitError int32  `json:"init_error"`
			InitErrno int32  `json:"init_errno"`
		} `json:"time"`
		Posnhd      [2][2]float32 `json:"posnhd"`
		ClassCount  int32         `json:"class_count"`
		Horaz       [30]float32   `json:"horaz"`
		Horel       [30]float32   `json:"horel"`
		McbDev      [64]byte      `json:"mcb_dev"`
		Hwid        byte          `json:"hwid"`
		padCgo1     [3]byte
		IwMotion    int32    `json:"iw_motion"`
		Lowtp       [2]int32 `json:"lowtp"`
		FormVersion int32    `json:"form_version"`
		Sterp       int32    `json:"sterp"`
		WrhdFs      [2]int32 `json:"wrhd_fs"`
		VfmXpnt     int32    `json:"vfm_xpnt"`
		Actual      struct {
			S2rec [2]struct {
				Rstate        int32 `json:"rstate"`
				RstateValid   int32 `json:"rstate_valid"`
				Position      int32 `json:"position"`
				Posvar        int32 `json:"posvar"`
				PositionValid int32 `json:"position_valid"`
			} `json:"s2rec"`
			S2recInuse int32 `json:"s2rec_inuse"`
		} `json:"actual"`
		Freqvc                [15]float32   `json:"freqvc"`
		Ibwvc                 [15]int32     `json:"ibwvc"`
		Ifp2vc                [16]int32     `json:"ifp2vc"`
		Cwrap                 [8]byte       `json:"cwrap"`
		Vacsw                 [2]int32      `json:"vacsw"`
		Motorv2               [2]float32    `json:"motorv2"`
		Itpthick2             [2]int32      `json:"itpthick2"`
		Thin                  [2]int32      `json:"thin"`
		Vac4                  [2]int32      `json:"vac4"`
		Wrvolt2               [2]float32    `json:"wrvolt2"`
		Wrvolt4               [2]float32    `json:"wrvolt4"`
		Wrvolt42              [2]float32    `json:"wrvolt42"`
		UserDev1Name          [2]byte       `json:"user_dev1_name"`
		UserDev2Name          [2]byte       `json:"user_dev2_name"`
		UserDev1Value         float64       `json:"user_dev1_value"`
		UserDev2Value         float64       `json:"user_dev2_value"`
		Rvac                  [2]Rvac_cmd   `json:"rvac"`
		Wvolt                 [2]Wvolt_cmd  `json:"wvolt"`
		Lo                    Lo_cmd        `json:"lo"`
		Pcalform              Pcalform_cmd  `json:"pcalform"`
		Pcald                 Pcald_cmd     `json:"pcald"`
		Extbwvc               [15]float32   `json:"extbwvc"`
		Freqif3               int32         `json:"freqif3"`
		Imixif3               int32         `json:"imixif3"`
		Pcalports             Pcalports_cmd `json:"pcalports"`
		K4RecState            int32         `json:"k4_rec_state"`
		K4st                  K4st_cmd      `json:"k4st"`
		K4tapeSqn             [9]byte       `json:"k4tape_sqn"`
		padCgo2               [3]byte
		K4vclo                K4vclo_cmd  `json:"k4vclo"`
		K4vc                  K4vc_cmd    `json:"k4vc"`
		K4vcif                K4vcif_cmd  `json:"k4vcif"`
		K4vcbw                K4vcbw_cmd  `json:"k4vcbw"`
		K3fm                  K3fm_cmd    `json:"k3fm"`
		Reccpu                [2]int32    `json:"reccpu"`
		K4label               K4label_cmd `json:"k4label"`
		padCgo3               [3]byte
		K4recMode             K4rec_mode_cmd  `json:"k4rec_mode"`
		K4recpatch            K4recpatch_cmd  `json:"k4recpatch"`
		K4pcalports           K4pcalports_cmd `json:"k4pcalports"`
		Select                int32           `json:"select"`
		RdhdFs                [2]int32        `json:"rdhd_fs"`
		Knewtape              [2]int32        `json:"knewtape"`
		Ihdmndel              [2]int32        `json:"ihdmndel"`
		ScanName              Scan_name_cmd   `json:"scan_name"`
		Tacd                  Tacd_shm        `json:"tacd"`
		Iat1if                int32           `json:"iat1if"`
		Iat2if                int32           `json:"iat2if"`
		Iat3if                int32           `json:"iat3if"`
		Erchk                 int32           `json:"erchk"`
		IfdSet                int32           `json:"ifd_set"`
		If3Set                int32           `json:"if3_set"`
		BbcTpi                [16][2]uint32   `json:"bbc_tpi"`
		VifdTpi               [4]uint32       `json:"vifd_tpi"`
		MifdTpi               [3]uint32       `json:"mifd_tpi"`
		Cablevl               float32         `json:"cablevl"`
		Cablediff             float32         `json:"cablediff"`
		Imk4fmv               int32           `json:"imk4fmv"`
		Tpicd                 Tpicd_cmd       `json:"tpicd"`
		ITPIVC                [15]int32       `json:"ITPIVC"`
		Tpigain               [264]int32      `json:"tpigain"`
		Iapdflg               int32           `json:"iapdflg"`
		K4recModeStat         int32           `json:"k4rec_mode_stat"`
		Onoff                 Onoff_cmd       `json:"onoff"`
		Rxgain                [20]Rxgain_ds   `json:"rxgain"`
		Iswif3Fs              [4]int32        `json:"iswif3_fs"`
		Ipcalif3              int32           `json:"ipcalif3"`
		Flux                  [100]Flux_ds    `json:"flux"`
		Tpidiff               [264]int32      `json:"tpidiff"`
		Tpidiffgain           [264]int32      `json:"tpidiffgain"`
		Caltemps              [264]float32    `json:"caltemps"`
		Calrx                 Calrx_cmd       `json:"calrx"`
		Ibds                  int32           `json:"ibds"`
		DsDev                 [64]byte        `json:"ds_dev"`
		NDas                  byte            `json:"n_das"`
		LbaImageRejectFilters byte            `json:"lba_image_reject_filters"`
		padCgo4               [2]byte
		LbaDigitalInputFormat uint32    `json:"lba_digital_input_format"`
		Das                   [2]Das    `json:"das"`
		IfpTpi                [4]uint32 `json:"ifp_tpi"`
		MDas                  byte      `json:"m_das"`
		Mk5vsn                [33]byte  `json:"mk5vsn"`
		padCgo5               [2]byte
		Mk5vsnLogchg          int32           `json:"mk5vsn_logchg"`
		Logchg                int32           `json:"logchg"`
		UserDevice            User_device_cmd `json:"user_device"`
		DiskRecord            Disk_record_cmd `json:"disk_record"`
		Monit5                struct {
			Pong int32          `json:"pong"`
			Ping [2]Monit5_ping `json:"ping"`
		} `json:"monit5"`
		Disk2file Disk2file_cmd `json:"disk2file"`
		In2net    In2net_cmd    `json:"in2net"`
		Abend     struct {
			NormalEnd  int32 `json:"normal_end"`
			OtherError int32 `json:"other_error"`
		} `json:"abend"`
		S2bbc           [4]S2bbc_data `json:"s2bbc"`
		S2das           S2das_check   `json:"s2das"`
		NtpSynchUnknown int32         `json:"ntp_synch_unknown"`
		LastCheck       struct {
			String  [256]byte `json:"string"`
			Ip2     int32     `json:"ip2"`
			Who     [3]byte   `json:"who"`
			padCgo0 [1]byte
		} `json:"last_check"`
		Mk5host        [129]byte `json:"mk5host"`
		padCgo6        [3]byte
		Mk5bMode       Mk5b_mode_cmd          `json:"mk5b_mode"`
		Vsi4           Vsi4_cmd               `json:"vsi4"`
		Holog          Holog_cmd              `json:"holog"`
		Satellite      Satellite_cmd          `json:"satellite"`
		Ephem          [14400]Satellite_ephem `json:"ephem"`
		Satoff         Satoff_cmd             `json:"satoff"`
		Tle            Tle_cmd                `json:"tle"`
		Dbbcnn         [16]Dbbcnn_cmd         `json:"dbbcnn"`
		Dbbcifx        [4]Dbbcifx_cmd         `json:"dbbcifx"`
		Dbbcform       Dbbcform_cmd           `json:"dbbcform"`
		Dbbcddcv       int32                  `json:"dbbcddcv"`
		Dbbcpfbv       int32                  `json:"dbbcpfbv"`
		DbbcCondMods   int32                  `json:"dbbc_cond_mods"`
		DbbcContCal    Dbbc_cont_cal_cmd      `json:"dbbc_cont_cal"`
		DbbcIfFactors  [4]int32               `json:"dbbc_if_factors"`
		Dbbcgain       Dbbcgain_cmd           `json:"dbbcgain"`
		M5bCrate       int32                  `json:"m5b_crate"`
		Dbbcddcvl      [1]byte                `json:"dbbcddcvl"`
		Dbbcddcvs      [16]byte               `json:"dbbcddcvs"`
		padCgo7        [3]byte
		Dbbcddcvc      int32            `json:"dbbcddcvc"`
		Dbbcddcsubv    int32            `json:"dbbcddcsubv"`
		Dbbccontcalpol int32            `json:"dbbccontcalpol"`
		Fila10gMode    Fila10g_mode_cmd `json:"fila10g_mode"`
		Fila10gvsiIn   [16]byte         `json:"fila10gvsi_in"`
		Dbbcpfbvl      [1]byte          `json:"dbbcpfbvl"`
		Dbbcpfbvs      [16]byte         `json:"dbbcpfbvs"`
		padCgo8        [3]byte
		Dbbcpfbvc      int32             `json:"dbbcpfbvc"`
		Dbbcpfbsubv    int32             `json:"dbbcpfbsubv"`
		DbbcComoCores  [4]int32          `json:"dbbc_como_cores"`
		DbbcCores      int32             `json:"dbbc_cores"`
		DbbcVsix       [2]Dbbc_vsix_cmd  `json:"dbbc_vsix"`
		Mk6Units       [2]int32          `json:"mk6_units"`
		Mk6Active      [2]int32          `json:"mk6_active"`
		Mk6Record      [3]Mk6_record_cmd `json:"mk6_record"`
		Mk6LastCheck   [2]struct {
			String  [256]byte `json:"string"`
			Ip2     int32     `json:"ip2"`
			Who     [3]byte   `json:"who"`
			What    [3]byte   `json:"what"`
			padCgo0 [2]byte
		} `json:"mk6_last_check"`
		RdbeUnits    [4]int32          `json:"rdbe_units"`
		RdbeActive   [4]int32          `json:"rdbe_active"`
		RdbeTsysData [4]Rdbe_tsys_data `json:"rdbe_tsys_data"`
		Rdbehost     [4][129]byte      `json:"rdbehost"`
		RdbeAtten    [5]Rdbe_atten_cmd `json:"rdbe_atten"`
		Rdtcn        [4]Rdtcn          `json:"rdtcn"`
		FserrCls     Fserr_cls         `json:"fserr_cls"`
		DbbcDefined  int32             `json:"dbbc_defined"`
		Dbbc2Defined int32             `json:"dbbc2_defined"`
		RdbeEquip    struct {
			RmsT    float32 `json:"rms_t"`
			RmsMin  float32 `json:"rms_min"`
			RmsMax  float32 `json:"rms_max"`
			PcalAmp [1]byte `json:"pcal_amp"`
			padCgo0 [3]byte
		} `json:"rdbe_equip"`
		Monit6            Monit6               `json:"monit6"`
		RdbeSync          [4]int32             `json:"rdbe_sync"`
		Dbbc3DdcV         int32                `json:"dbbc3_ddc_v"`
		Dbbc3DdcVs        [16]byte             `json:"dbbc3_ddc_vs"`
		Dbbc3DdcVc        int32                `json:"dbbc3_ddc_vc"`
		Dbbc3DdcBbcsPerIf int32                `json:"dbbc3_ddc_bbcs_per_if"`
		Dbbc3DdcIfs       int32                `json:"dbbc3_ddc_ifs"`
		Dbbc3Ifx          [8]Dbbc3_ifx_cmd     `json:"dbbc3_ifx"`
		Dbbc3Bbcnn        [128]Dbbc3_bbcnn_cmd `json:"dbbc3_bbcnn"`
		Dbbc3ContCal      Dbbc3_cont_cal_cmd   `json:"dbbc3_cont_cal"`
		SVerReleaseFS     [33]byte             `json:"sVerRelease_FS"`
		padCgo9           [3]byte
	}
	Fserr_cls struct {
		Buf     [125]byte `json:"buf"`
		padCgo0 [3]byte
		Nchars  int32 `json:"nchars"`
	}
	Ft struct {
		Sync            uint32 `json:"sync"`
		NcoCentreValue  uint32 `json:"nco_centre_value"`
		NcoOffsetValue  uint32 `json:"nco_offset_value"`
		NcoPhaseValue   uint32 `json:"nco_phase_value"`
		NcoTimerValue   uint32 `json:"nco_timer_value"`
		NcoTest         uint32 `json:"nco_test"`
		NcoUseOffset    uint32 `json:"nco_use_offset"`
		NcoSyncReset    uint32 `json:"nco_sync_reset"`
		NcoUseTimer     uint32 `json:"nco_use_timer"`
		QFirNo          byte   `json:"q_fir_no"`
		IFirNo          byte   `json:"i_fir_no"`
		ClockDecimation byte   `json:"clock_decimation"`
		padCgo0         [1]byte
		AddSub          Mux  `json:"add_sub"`
		UsbMux          Mux  `json:"usb_mux"`
		LsbMux          Mux  `json:"lsb_mux"`
		UsbThreshold    byte `json:"usb_threshold"`
		LsbThreshold    byte `json:"lsb_threshold"`
		padCgo1         [2]byte
		UsbServo        Servo  `json:"usb_servo"`
		LsbServo        Servo  `json:"lsb_servo"`
		FlipUsb         uint32 `json:"flip_usb"`
		FlipLsb         uint32 `json:"flip_lsb"`
		Monitor         Mux    `json:"monitor"`
		Digout          Digout `json:"digout"`
	}
	Holog_cmd struct {
		Az          float32  `json:"az"`
		El          float32  `json:"el"`
		Azp         int32    `json:"azp"`
		Elp         int32    `json:"elp"`
		Ical        int32    `json:"ical"`
		Proc        [33]byte `json:"proc"`
		padCgo0     [3]byte
		StopRequest int32 `json:"stop_request"`
		Setup       int32 `json:"setup"`
		Wait        int32 `json:"wait"`
	}
	Ifp struct {
		Frequency    float64   `json:"frequency"`
		Bandwidth    uint32    `json:"bandwidth"`
		FilterMode   uint32    `json:"filter_mode"`
		FlipUsb      uint32    `json:"flip_usb"`
		FlipLsb      uint32    `json:"flip_lsb"`
		Format       uint32    `json:"format"`
		MagnStats    uint32    `json:"magn_stats"`
		CorrType     uint32    `json:"corr_type"`
		CorrSource   [2]uint32 `json:"corr_source"`
		AtClockDelay byte      `json:"at_clock_delay"`
		padCgo0      [3]byte
		FtLo         float64 `json:"ft_lo"`
		FtFilterMode uint32  `json:"ft_filter_mode"`
		FtOffs       float64 `json:"ft_offs"`
		FtPhase      float64 `json:"ft_phase"`
		Track        [2]byte `json:"track"`
		Initialised  byte    `json:"initialised"`
		padCgo1      [1]byte
		Source       int32   `json:"source"`
		FilterOutput uint32  `json:"filter_output"`
		Bs           Bs      `json:"bs"`
		Ft           Ft      `json:"ft"`
		Out          Out     `json:"out"`
		TempAnalog   float32 `json:"temp_analog"`
		PllLd        float32 `json:"pll_ld"`
		PllVc        float32 `json:"pll_vc"`
		RefErr       byte    `json:"ref_err"`
		SyncErr      byte    `json:"sync_err"`
		padCgo2      [2]byte
		TempDigital  float32 `json:"temp_digital"`
		Processing   byte    `json:"processing"`
		ClkErr       byte    `json:"clk_err"`
		Blank        byte    `json:"blank"`
		padCgo3      [1]byte
	}
	In2net_cmd struct {
		Control struct {
			Control int32   `json:"control"`
			State   M5state `json:"state"`
		} `json:"control"`
		Destination struct {
			Destination [33]byte `json:"destination"`
			padCgo0     [3]byte
			State       M5state `json:"state"`
		} `json:"destination"`
		Options struct {
			Options [33]byte `json:"options"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"options"`
		LastDestination [33]byte `json:"last_destination"`
		padCgo0         [3]byte
	}
	In2net_mon struct {
		Status struct {
			Status  [33]byte `json:"status"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"status"`
		Received struct {
			Received int64   `json:"received"`
			State    M5state `json:"state"`
		} `json:"received"`
		Buffered struct {
			Buffered int64   `json:"buffered"`
			State    M5state `json:"state"`
		} `json:"buffered"`
	}
	K3fm_cmd struct {
		Mode     int32    `json:"mode"`
		Rate     int32    `json:"rate"`
		Input    int32    `json:"input"`
		Aux      [12]byte `json:"aux"`
		Synch    int32    `json:"synch"`
		AuxStart int32    `json:"aux_start"`
		Output   int32    `json:"output"`
	}
	K3fm_mon struct {
		Daytime [15]byte `json:"daytime"`
		Status  [3]byte  `json:"status"`
	}
	K4label_cmd struct {
		Label [9]byte `json:"label"`
	}
	K4pcalports_cmd struct {
		Ports [2]int32 `json:"ports"`
	}
	K4pcalports_mon struct {
		Amp   [2]float32 `json:"amp"`
		Phase [2]float32 `json:"phase"`
	}
	K4rec_check struct {
		Check int32 `json:"check"`
		State int32 `json:"state"`
		Mode  int32 `json:"mode"`
		Pca   int32 `json:"pca"`
		Pcb   int32 `json:"pcb"`
		Drm   int32 `json:"drm"`
		Synch int32 `json:"synch"`
		Aux   int32 `json:"aux"`
	}
	K4rec_mode_cmd struct {
		Bw int32 `json:"bw"`
		Bt int32 `json:"bt"`
		Ch int32 `json:"ch"`
		Im int32 `json:"im"`
		Nm int32 `json:"nm"`
	}
	K4rec_mode_mon struct {
		Ts int32 `json:"ts"`
		Fm int32 `json:"fm"`
		Ta int32 `json:"ta"`
		Pb int32 `json:"pb"`
	}
	K4recpatch_cmd struct {
		Ports [16]int32 `json:"ports"`
	}
	K4st_cmd struct {
		Record int32 `json:"record"`
	}
	K4vcbw_cmd struct {
		Bw [2]int32 `json:"bw"`
	}
	K4vc_cmd struct {
		Lohi [16]int32 `json:"lohi"`
		Att  [16]int32 `json:"att"`
		Loup [16]int32 `json:"loup"`
	}
	K4vcif_cmd struct {
		Att [4]int32 `json:"att"`
	}
	K4vclo_cmd struct {
		Freq [16]int32 `json:"freq"`
	}
	K4vclo_mon struct {
		Yes  [16]byte `json:"yes"`
		Lock [16]byte `json:"lock"`
	}
	K4vc_mon struct {
		Yes    [16]byte  `json:"yes"`
		Usbpwr [16]int32 `json:"usbpwr"`
		Lsbpwr [16]int32 `json:"lsbpwr"`
	}
	Lo_cmd struct {
		Lo       [8]float64 `json:"lo"`
		Sideband [8]int32   `json:"sideband"`
		Pol      [8]int32   `json:"pol"`
		Spacing  [8]float64 `json:"spacing"`
		Offset   [8]float64 `json:"offset"`
		Pcal     [8]int32   `json:"pcal"`
	}
	M5state struct {
		Known int32 `json:"known"`
		Error int32 `json:"error"`
	}
	M5time struct {
		Year             int32   `json:"year"`
		Day              int32   `json:"day"`
		Hour             int32   `json:"hour"`
		Minute           int32   `json:"minute"`
		Seconds          float64 `json:"seconds"`
		SecondsPrecision int32   `json:"seconds_precision"`
	}
	Mcb_cmd struct {
		Device  [2]byte `json:"device"`
		padCgo0 [2]byte
		Addr    uint32 `json:"addr"`
		Data    uint32 `json:"data"`
		Cmd     int32  `json:"cmd"`
	}
	Mcb_mon struct {
		Data uint32 `json:"data"`
	}
	Mk5b_mode_cmd struct {
		Source struct {
			Source  int32    `json:"source"`
			Magic   [33]byte `json:"magic"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"source"`
		Mask struct {
			Mask  uint64  `json:"mask"`
			Bits  int32   `json:"bits"`
			State M5state `json:"state"`
		} `json:"mask"`
		Decimate struct {
			Decimate int32   `json:"decimate"`
			Datarate uint64  `json:"datarate"`
			State    M5state `json:"state"`
		} `json:"decimate"`
		Samplerate struct {
			Samplerate uint64  `json:"samplerate"`
			Decimate   int32   `json:"decimate"`
			Datarate   uint64  `json:"datarate"`
			State      M5state `json:"state"`
		} `json:"samplerate"`
		Fpdp struct {
			Fpdp  int32   `json:"fpdp"`
			State M5state `json:"state"`
		} `json:"fpdp"`
		Disk struct {
			Disk  int32   `json:"disk"`
			State M5state `json:"state"`
		} `json:"disk"`
	}
	Mk5b_mode_mon struct {
		Format struct {
			Format  [33]byte `json:"format"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"format"`
		Tracks struct {
			Tracks int32   `json:"tracks"`
			State  M5state `json:"state"`
		} `json:"tracks"`
		Tbitrate struct {
			Tbitrate float64 `json:"tbitrate"`
			State    M5state `json:"state"`
		} `json:"tbitrate"`
		Framesize struct {
			Framesize int32   `json:"framesize"`
			State     M5state `json:"state"`
		} `json:"framesize"`
	}
	Mk6_disk_pos_mon struct {
		Record struct {
			Record int64   `json:"record"`
			State  M5state `json:"state"`
		} `json:"record"`
		Play struct {
			Play  int64   `json:"play"`
			State M5state `json:"state"`
		} `json:"play"`
		Stop struct {
			Stop  int64   `json:"stop"`
			State M5state `json:"state"`
		} `json:"stop"`
	}
	Mk6_record_cmd struct {
		Action struct {
			Action  [22]byte `json:"action"`
			padCgo0 [2]byte
			State   M5state `json:"state"`
		} `json:"action"`
		Duration struct {
			Duration int32   `json:"duration"`
			State    M5state `json:"state"`
		} `json:"duration"`
		Size struct {
			Size  int32   `json:"size"`
			State M5state `json:"state"`
		} `json:"size"`
		Scan struct {
			Scan    [33]byte `json:"scan"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"scan"`
		Experiment struct {
			Experiment [9]byte `json:"experiment"`
			padCgo0    [3]byte
			State      M5state `json:"state"`
		} `json:"experiment"`
		Station struct {
			Station [9]byte `json:"station"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"station"`
	}
	Mk6_record_mon struct {
		Status struct {
			Status  [33]byte `json:"status"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"status"`
		Group struct {
			Group int32   `json:"group"`
			State M5state `json:"state"`
		} `json:"group"`
		Number struct {
			Number int32   `json:"number"`
			State  M5state `json:"state"`
		} `json:"number"`
		Name struct {
			Name    [33]byte `json:"name"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"name"`
	}
	Mk6_scan_check_mon struct {
		Scan struct {
			Scan  int32   `json:"scan"`
			State M5state `json:"state"`
		} `json:"scan"`
		Label struct {
			Label   [65]byte `json:"label"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"label"`
		Type struct {
			Type    [33]byte `json:"type"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"type"`
		Code struct {
			Code  int32   `json:"code"`
			State M5state `json:"state"`
		} `json:"code"`
		Start struct {
			Start M5time  `json:"start"`
			State M5state `json:"state"`
		} `json:"start"`
		Length struct {
			Length M5time  `json:"length"`
			State  M5state `json:"state"`
		} `json:"length"`
		Total struct {
			Total float32 `json:"total"`
			State M5state `json:"state"`
		} `json:"total"`
		Missing struct {
			Missing int64   `json:"missing"`
			State   M5state `json:"state"`
		} `json:"missing"`
		Error struct {
			Error   [33]byte `json:"error"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"error"`
	}
	Monit5_ping struct {
		Active int32 `json:"active"`
		Bank   [2]struct {
			Vsn     [33]byte `json:"vsn"`
			padCgo0 [3]byte
			Seconds float64  `json:"seconds"`
			Gb      float64  `json:"gb"`
			Percent float64  `json:"percent"`
			Itime   [6]int32 `json:"itime"`
		} `json:"bank"`
	}
	Monit6 struct {
		Tsys      [2][4]int32 `json:"tsys"`
		Pcal      [2][4]int32 `json:"pcal"`
		Dot2ppsNs int32       `json:"dot2pps_ns"`
	}
	Mux struct {
		Setting byte `json:"setting"`
		padCgo0 [3]byte
		Mode    uint32 `json:"mode"`
	}
	Onoff_cmd struct {
		Rep         int32    `json:"rep"`
		Intp        int32    `json:"intp"`
		Cutoff      float32  `json:"cutoff"`
		Step        float32  `json:"step"`
		Wait        int32    `json:"wait"`
		Ssize       float32  `json:"ssize"`
		Proc        [33]byte `json:"proc"`
		padCgo0     [3]byte
		Devices     [270]Onoff_devices `json:"devices"`
		Itpis       [270]int32         `json:"itpis"`
		Fwhm        float32            `json:"fwhm"`
		StopRequest int32              `json:"stop_request"`
		Setup       int32              `json:"setup"`
	}
	Onoff_devices struct {
		Lwhat   [4]byte `json:"lwhat"`
		Pol     byte    `json:"pol"`
		padCgo0 [3]byte
		Ifchain int32   `json:"ifchain"`
		Flux    float32 `json:"flux"`
		Corr    float32 `json:"corr"`
		Center  float64 `json:"center"`
		Fwhm    float32 `json:"fwhm"`
		Tcal    float32 `json:"tcal"`
		Dpfu    float32 `json:"dpfu"`
		Gain    float32 `json:"gain"`
	}
	Out struct {
		S2Lo           S2_out `json:"s2_lo"`
		S2Hi           S2_out `json:"s2_hi"`
		AtmbCorrSource uint32 `json:"atmb_corr_source"`
		MbCorr2Source  uint32 `json:"mb_corr_2_source"`
		AtClockDelay   byte   `json:"at_clock_delay"`
		padCgo0        [3]byte
	}
	Pcald_cmd struct {
		Continuous  int32              `json:"continuous"`
		Bits        int32              `json:"bits"`
		Integration int32              `json:"integration"`
		StopRequest int32              `json:"stop_request"`
		Count       [2][16]int32       `json:"count"`
		Freqs       [2][16][17]float64 `json:"freqs"`
	}
	Pcalform_cmd struct {
		Count  [2][16]int32       `json:"count"`
		Which  [2][16][17]int32   `json:"which"`
		Tones  [2][16][17]int32   `json:"tones"`
		Strlen [2][16][17]int32   `json:"strlen"`
		Freqs  [2][16][17]float64 `json:"freqs"`
	}
	Pcalports_cmd struct {
		Bbc [2]int32 `json:"bbc"`
	}
	Pps_source_cmd struct {
		Source struct {
			Source  [33]byte `json:"source"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"source"`
	}
	Rclcn_req_buf struct {
		Count      int32     `json:"count"`
		ClassFs    int32     `json:"class_fs"`
		Nchars     int32     `json:"nchars"`
		PrevNchars int32     `json:"prev_nchars"`
		Buf        [512]byte `json:"buf"`
	}
	Rclcn_res_buf struct {
		ClassFs int32     `json:"class_fs"`
		Count   int32     `json:"count"`
		Ifc     int32     `json:"ifc"`
		Nchars  int32     `json:"nchars"`
		Buf     [512]byte `json:"buf"`
	}
	Rdbe_atten_cmd struct {
		Ifc struct {
			Ifc   int32   `json:"ifc"`
			State M5state `json:"state"`
		} `json:"ifc"`
		Atten struct {
			Atten int32   `json:"atten"`
			State M5state `json:"state"`
		} `json:"atten"`
		Target struct {
			Target float32 `json:"target"`
			State  M5state `json:"state"`
		} `json:"target"`
	}
	Rdbe_atten_mon struct {
		Ifc [2]struct {
			Ifc struct {
				Ifc   int32   `json:"ifc"`
				State M5state `json:"state"`
			} `json:"ifc"`
			Atten struct {
				Atten int32   `json:"atten"`
				State M5state `json:"state"`
			} `json:"atten"`
			RMS struct {
				RMS   float32 `json:"RMS"`
				State M5state `json:"state"`
			} `json:"RMS"`
		} `json:"ifc"`
	}
	Rdbe_dot_mon struct {
		Time struct {
			Time    [33]byte `json:"time"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"time"`
		Status struct {
			Status  [33]byte `json:"status"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"status"`
		OSTime struct {
			OSTime  [33]byte `json:"OS_time"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"OS_time"`
		DOTOSTimeDiff struct {
			DOTOSTimeDiff [33]byte `json:"DOT_OS_time_diff"`
			padCgo0       [3]byte
			State         M5state `json:"state"`
		} `json:"DOT_OS_time_diff"`
		ActualDOTTime struct {
			ActualDOTTime [33]byte `json:"Actual_DOT_time"`
			padCgo0       [3]byte
			State         M5state `json:"state"`
		} `json:"Actual_DOT_time"`
		VdifEpoch struct {
			VdifEpoch int32   `json:"vdif_epoch"`
			State     M5state `json:"state"`
		} `json:"vdif_epoch"`
	}
	Rdbe_tsys_cycle struct {
		Epoch       [14]byte `json:"epoch"`
		padCgo0     [2]byte
		EpochVdif   int32          `json:"epoch_vdif"`
		Tsys        [18][2]float32 `json:"tsys"`
		PcalAmp     [512]float32   `json:"pcal_amp"`
		PcalPhase   [512]float32   `json:"pcal_phase"`
		PcalIfx     int32          `json:"pcal_ifx"`
		Sigma       float32        `json:"sigma"`
		RawIfx      int32          `json:"raw_ifx"`
		Dot2gps     float64        `json:"dot2gps"`
		Dot2pps     float64        `json:"dot2pps"`
		Pcaloff     float64        `json:"pcaloff"`
		PcalSpacing float64        `json:"pcal_spacing"`
	}
	Rdbe_tsys_data struct {
		Data  [2]Rdbe_tsys_cycle `json:"data"`
		Iping int32              `json:"iping"`
	}
	Rdtcn struct {
		Control [2]Rdtcn_control `json:"control"`
		Iping   int32            `json:"iping"`
	}
	Rdtcn_control struct {
		Continuous  int32          `json:"continuous"`
		Cycle       int32          `json:"cycle"`
		StopRequest int32          `json:"stop_request"`
		DataValid   Data_valid_cmd `json:"data_valid"`
	}
	Rec_mode_cmd struct {
		Mode      [21]byte `json:"mode"`
		padCgo0   [3]byte
		Group     int32 `json:"group"`
		Roll      int32 `json:"roll"`
		NumGroups int32 `json:"num_groups"`
	}
	Regs struct {
		Error   byte `json:"error"`
		Warning byte `json:"warning"`
	}
	Req_buf struct {
		Count   int32     `json:"count"`
		ClassFs int32     `json:"class_fs"`
		Nchars  int32     `json:"nchars"`
		Buf     [512]byte `json:"buf"`
	}
	Req_rec struct {
		Type    int32   `json:"type"`
		Device  [2]byte `json:"device"`
		padCgo0 [2]byte
		Addr    uint32 `json:"addr"`
		Data    uint32 `json:"data"`
	}
	Res_buf struct {
		ClassFs int32     `json:"class_fs"`
		Count   int32     `json:"count"`
		Ifc     int32     `json:"ifc"`
		Nchars  int32     `json:"nchars"`
		Buf     [512]byte `json:"buf"`
	}
	Res_rec struct {
		State int32    `json:"state"`
		Code  int32    `json:"code"`
		Data  uint32   `json:"data"`
		Array [24]byte `json:"array"`
	}
	Rtime_mon struct {
		Seconds struct {
			Seconds float64 `json:"seconds"`
			State   M5state `json:"state"`
		} `json:"seconds"`
		Gb struct {
			Gb    float64 `json:"gb"`
			State M5state `json:"state"`
		} `json:"gb"`
		Percent struct {
			Percent float64 `json:"percent"`
			State   M5state `json:"state"`
		} `json:"percent"`
		TotalRate struct {
			TotalRate float64 `json:"total_rate"`
			State     M5state `json:"state"`
		} `json:"total_rate"`
		Mode struct {
			Mode    [33]byte `json:"mode"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"mode"`
		SubMode struct {
			SubMode [33]byte `json:"sub_mode"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"sub_mode"`
		TrackRate struct {
			TrackRate float64 `json:"track_rate"`
			State     M5state `json:"state"`
		} `json:"track_rate"`
		Source struct {
			Source  [33]byte `json:"source"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"source"`
		Mask struct {
			Mask  uint32  `json:"mask"`
			State M5state `json:"state"`
		} `json:"mask"`
		Decimate struct {
			Decimate int32   `json:"decimate"`
			State    M5state `json:"state"`
		} `json:"decimate"`
	}
	Rvac_cmd struct {
		Inches float32 `json:"inches"`
		Set    int32   `json:"set"`
	}
	Rvac_mon struct {
		Volts float32 `json:"volts"`
	}
	Rxgain_ds struct {
		Type    byte `json:"type"`
		padCgo0 [3]byte
		Lo      [2]float32 `json:"lo"`
		Year    int32      `json:"year"`
		Month   int32      `json:"month"`
		Day     int32      `json:"day"`
		Fwhm    struct {
			Model   byte `json:"model"`
			padCgo0 [3]byte
			Coeff   float32 `json:"coeff"`
		} `json:"fwhm"`
		Pol     [2]byte `json:"pol"`
		padCgo1 [2]byte
		Dpfu    [2]float32 `json:"dpfu"`
		Gain    struct {
			Form    byte `json:"form"`
			Type    byte `json:"type"`
			padCgo0 [2]byte
			Coeff   [10]float32 `json:"coeff"`
			Ncoeff  int32       `json:"ncoeff"`
			Opacity byte        `json:"opacity"`
			padCgo1 [3]byte
		} `json:"gain"`
		TcalNtable int32    `json:"tcal_ntable"`
		TcalNpol   [2]int32 `json:"tcal_npol"`
		Tcal       [1200]struct {
			Pol     byte `json:"pol"`
			padCgo0 [3]byte
			Freq    float32 `json:"freq"`
			Tcal    float32 `json:"tcal"`
		} `json:"tcal"`
		Trec        [2]float32 `json:"trec"`
		SpillNtable int32      `json:"spill_ntable"`
		Spill       [20]struct {
			El float32 `json:"el"`
			Tk float32 `json:"tk"`
		} `json:"spill"`
	}
	S2bbc_data struct {
		Freq    uint32  `json:"freq"`
		Tpiavg  uint16  `json:"tpiavg"`
		Ifsrc   byte    `json:"ifsrc"`
		Bw      [2]byte `json:"bw"`
		Agcmode byte    `json:"agcmode"`
		Init    byte    `json:"init"`
		padCgo0 [1]byte
	}
	S2das_check struct {
		Check    uint32   `json:"check"`
		Agc      byte     `json:"agc"`
		Encode   byte     `json:"encode"`
		Mode     [21]byte `json:"mode"`
		FSstatus byte     `json:"FSstatus"`
		SeqName  [25]byte `json:"SeqName"`
		BW       byte     `json:"BW"`
		padCgo0  [2]byte
	}
	S2label_cmd struct {
		Tapeid   [21]byte `json:"tapeid"`
		Tapetype [7]byte  `json:"tapetype"`
		Format   [33]byte `json:"format"`
	}
	S2_out struct {
		Source uint32 `json:"source"`
		Format uint32 `json:"format"`
	}
	S2rec_check struct {
		Check    int32 `json:"check"`
		UserInfo struct {
			Label [4]int32 `json:"label"`
			Field [4]int32 `json:"field"`
		} `json:"user_info"`
		Speed    int32 `json:"speed"`
		State    int32 `json:"state"`
		Mode     int32 `json:"mode"`
		Group    int32 `json:"group"`
		Roll     int32 `json:"roll"`
		Dv       int32 `json:"dv"`
		Tapeid   int32 `json:"tapeid"`
		Tapetype int32 `json:"tapetype"`
	}
	S2st_cmd struct {
		Dir    int32 `json:"dir"`
		Speed  int32 `json:"speed"`
		Record int32 `json:"record"`
	}
	Satellite_cmd struct {
		Name      [25]byte `json:"name"`
		Tlefile   [65]byte `json:"tlefile"`
		padCgo0   [2]byte
		Mode      int32    `json:"mode"`
		Wrap      int32    `json:"wrap"`
		Satellite int32    `json:"satellite"`
		Tle0      [25]byte `json:"tle0"`
		Tle1      [70]byte `json:"tle1"`
		Tle2      [70]byte `json:"tle2"`
		padCgo1   [3]byte
	}
	Satellite_ephem struct {
		T  int32   `json:"t"`
		Az float64 `json:"az"`
		El float64 `json:"el"`
	}
	Satoff_cmd struct {
		Seconds float64 `json:"seconds"`
		Cross   float64 `json:"cross"`
		Hold    int32   `json:"hold"`
	}
	Scan_check_mon struct {
		Scan struct {
			Scan  int32   `json:"scan"`
			State M5state `json:"state"`
		} `json:"scan"`
		Label struct {
			Label   [65]byte `json:"label"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"label"`
		Start struct {
			Start M5time  `json:"start"`
			State M5state `json:"state"`
		} `json:"start"`
		Length struct {
			Length M5time  `json:"length"`
			State  M5state `json:"state"`
		} `json:"length"`
		Missing struct {
			Missing int64   `json:"missing"`
			State   M5state `json:"state"`
		} `json:"missing"`
		Mode struct {
			Mode    [33]byte `json:"mode"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"mode"`
		Submode struct {
			Submode [33]byte `json:"submode"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"submode"`
		Rate struct {
			Rate  float32 `json:"rate"`
			State M5state `json:"state"`
		} `json:"rate"`
		Type struct {
			Type    [33]byte `json:"type"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"type"`
		Code struct {
			Code  int32   `json:"code"`
			State M5state `json:"state"`
		} `json:"code"`
		Total struct {
			Total float32 `json:"total"`
			State M5state `json:"state"`
		} `json:"total"`
		Error struct {
			Error   [33]byte `json:"error"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"error"`
	}
	Scan_name_cmd struct {
		NameOld    [17]byte `json:"name_old"`
		Name       [17]byte `json:"name"`
		Session    [17]byte `json:"session"`
		Station    [17]byte `json:"station"`
		Duration   int32    `json:"duration"`
		Continuous int32    `json:"continuous"`
	}
	Servo struct {
		Setting uint16 `json:"setting"`
		padCgo0 [2]byte
		Mode    uint32 `json:"mode"`
		Readout int32  `json:"readout"`
	}
	Systracks_cmd struct {
		Track [4]int32 `json:"track"`
	}
	Tacd_shm struct {
		Day               int32    `json:"day"`
		DayFrac           int32    `json:"day_frac"`
		MsecCounter       float32  `json:"msec_counter"`
		UsecCorrection    float32  `json:"usec_correction"`
		UsecBias          float32  `json:"usec_bias"`
		CookedCorrection  float32  `json:"cooked_correction"`
		PcVUtc            float32  `json:"pc_v_utc"`
		UtcCorrectionNsec float32  `json:"utc_correction_nsec"`
		UtcCorrectionSec  int32    `json:"utc_correction_sec"`
		DayA              int32    `json:"day_a"`
		DayFracA          int32    `json:"day_frac_a"`
		Rms               float32  `json:"rms"`
		UsecAverage       float32  `json:"usec_average"`
		Max               float32  `json:"max"`
		Min               float32  `json:"min"`
		DayFracOld        int32    `json:"day_frac_old"`
		DayFracOldA       int32    `json:"day_frac_old_a"`
		Continuous        int32    `json:"continuous"`
		NsecAccuracy      int32    `json:"nsec_accuracy"`
		SecAverage        int32    `json:"sec_average"`
		StopRequest       int32    `json:"stop_request"`
		Port              int32    `json:"port"`
		Check             int32    `json:"check"`
		Display           int32    `json:"display"`
		Hostpc            [80]byte `json:"hostpc"`
		Oldnew            [8]byte  `json:"oldnew"`
		OldnewA           [11]byte `json:"oldnew_a"`
		File              [40]byte `json:"file"`
		Status            [8]byte  `json:"status"`
		TacVer            [20]byte `json:"tac_ver"`
		padCgo0           [1]byte
	}
	Tape_cmd struct {
		Set   int32 `json:"set"`
		Reset int32 `json:"reset"`
	}
	Tape_mon struct {
		Foot    int32 `json:"foot"`
		Sense   int32 `json:"sense"`
		Vacuum  int32 `json:"vacuum"`
		Chassis int32 `json:"chassis"`
		Stat    int32 `json:"stat"`
		Error   int32 `json:"error"`
	}
	Tle_cmd struct {
		Tle0    [25]byte `json:"tle0"`
		Tle1    [70]byte `json:"tle1"`
		Tle2    [70]byte `json:"tle2"`
		padCgo0 [3]byte
		Catnum  [3]int32 `json:"catnum"`
	}
	Tpicd_cmd struct {
		Continuous  int32        `json:"continuous"`
		Cycle       int32        `json:"cycle"`
		StopRequest int32        `json:"stop_request"`
		Itpis       [264]int32   `json:"itpis"`
		Ifc         [264]int32   `json:"ifc"`
		Lwhat       [264][4]byte `json:"lwhat"`
		TsysRequest int32        `json:"tsys_request"`
	}
	User_device_cmd struct {
		Lo       [6]float64 `json:"lo"`
		Sideband [6]int32   `json:"sideband"`
		Pol      [6]int32   `json:"pol"`
		Center   [6]float64 `json:"center"`
		Zero     [6]int32   `json:"zero"`
	}
	User_info_cmd struct {
		Labels [4][17]byte `json:"labels"`
		Field1 [17]byte    `json:"field1"`
		Field2 [17]byte    `json:"field2"`
		Field3 [33]byte    `json:"field3"`
		Field4 [49]byte    `json:"field4"`
	}
	User_info_parse struct {
		Field   int32    `json:"field"`
		Label   int32    `json:"label"`
		String  [49]byte `json:"string"`
		padCgo0 [3]byte
	}
	Venable_cmd struct {
		General int32    `json:"general"`
		Group   [8]int32 `json:"group"`
	}
	Vform_cmd struct {
		Mode   int32 `json:"mode"`
		Rate   int32 `json:"rate"`
		Format int32 `json:"format"`
		Enable struct {
			Low    uint32 `json:"low"`
			High   uint32 `json:"high"`
			System uint32 `json:"system"`
		} `json:"enable"`
		Aux       [28][4]uint32 `json:"aux"`
		Codes     [32]int32     `json:"codes"`
		Fan       int32         `json:"fan"`
		Barrel    int32         `json:"barrel"`
		TapeClock int32         `json:"tape_clock"`
		Qa        struct {
			Drive int32 `json:"drive"`
			Chan  int32 `json:"chan"`
		} `json:"qa"`
		Last int32 `json:"last"`
	}
	Vform_mon struct {
		Version int32 `json:"version"`
		SysSt   int32 `json:"sys_st"`
		McbSt   int32 `json:"mcb_st"`
		HdwSt   int32 `json:"hdw_st"`
		SfwSt   int32 `json:"sfw_st"`
		IntSt   int32 `json:"int_st"`
	}
	Vrepro_cmd struct {
		Mode      [2]int32 `json:"mode"`
		Track     [2]int32 `json:"track"`
		Head      [2]int32 `json:"head"`
		Equalizer [2]int32 `json:"equalizer"`
		Bitsynch  int32    `json:"bitsynch"`
	}
	Vsi4 struct {
		Value int32 `json:"value"`
		Set   int32 `json:"set"`
	}
	Vsi4_cmd struct {
		Config Vsi4 `json:"config"`
		Pcalx  Vsi4 `json:"pcalx"`
		Pcaly  Vsi4 `json:"pcaly"`
	}
	Vsi4_mon struct {
		Version int32 `json:"version"`
	}
	Vsn_mon struct {
		Vsn struct {
			Vsn     [33]byte `json:"vsn"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"vsn"`
		Check struct {
			Check   [33]byte `json:"check"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"check"`
		Disk struct {
			Disk  int32   `json:"disk"`
			State M5state `json:"state"`
		} `json:"disk"`
		OriginalVsn struct {
			OriginalVsn [33]byte `json:"original_vsn"`
			padCgo0     [3]byte
			State       M5state `json:"state"`
		} `json:"original_vsn"`
		NewVsn struct {
			NewVsn  [33]byte `json:"new_vsn"`
			padCgo0 [3]byte
			State   M5state `json:"state"`
		} `json:"new_vsn"`
	}
	Vst_cmd struct {
		Dir   int32  `json:"dir"`
		Speed int32  `json:"speed"`
		Cips  uint32 `json:"cips"`
		Rec   int32  `json:"rec"`
	}
	Wvolt_cmd struct {
		Volts [2]float32 `json:"volts"`
		Set   [2]int32   `json:"set"`
	}
)
