// +build ignore
// generated with types.sh --- do not edit

package fs

/*
#include "/usr2/fs/include/params.h"
#include "/usr2/fs/include/fs_types.h"
#include "/usr2/fs/include/fscom.h"
*/
import "C"

type (
	Bank_set_mon      C.struct_bank_set_mon
	Bbc_cmd           C.struct_bbc_cmd
	Bbc_mon           C.struct_bbc_mon
	Bs                C.struct_bs
	Calrx_cmd         C.struct_calrx_cmd
	Capture_mon       C.struct_capture_mon
	Clock_set_cmd     C.struct_clock_set_cmd
	Cmd_ds            C.struct_cmd_ds
	Das               C.struct_das
	Data_check_mon    C.struct_data_check_mon
	Data_valid_cmd    C.struct_data_valid_cmd
	Dbbc_cont_cal_cmd C.struct_dbbc_cont_cal_cmd
	Dbbcform_cmd      C.struct_dbbcform_cmd
	Dbbcgain_cmd      C.struct_dbbcgain_cmd
	Dbbcgain_mon      C.struct_dbbcgain_mon
	Dbbcifx_cmd       C.struct_dbbcifx_cmd
	Dbbcifx_mon       C.struct_dbbcifx_mon
	Dbbcnn_cmd        C.struct_dbbcnn_cmd
	Dbbcnn_mon        C.struct_dbbcnn_mon
	Dbbc_pfbx_mon     C.struct_dbbc_pfbx_mon
	Dbbcvsi_clk_mon   C.struct_dbbcvsi_clk_mon
	Dbbc_vsix_cmd     C.struct_dbbc_vsix_cmd
	Digout            C.struct_digout
	Disk2file_cmd     C.struct_disk2file_cmd
	Disk2file_mon     C.struct_disk2file_mon
	Disk_pos_mon      C.struct_disk_pos_mon
	Disk_record_cmd   C.struct_disk_record_cmd
	Disk_record_mon   C.struct_disk_record_mon
	Disk_serial_mon   C.struct_disk_serial_mon
	Dist_cmd          C.struct_dist_cmd
	Dist_mon          C.struct_dist_mon
	Dot_mon           C.struct_dot_mon
	Dqa_chan          C.struct_dqa_chan
	Dqa_cmd           C.struct_dqa_cmd
	Dqa_mon           C.struct_dqa_mon
	Ds_cmd            C.struct_ds_cmd
	Ds_mon            C.struct_ds_mon
	Fila10g_mode_cmd  C.struct_fila10g_mode_cmd
	Fila10g_mode_mon  C.struct_fila10g_mode_mon
	Flux_ds           C.struct_flux_ds
	Form4_cmd         C.struct_form4_cmd
	Form4_mon         C.struct_form4_mon
	Fscom             C.struct_fscom
	Fserr_cls         C.struct_fserr_cls
	Ft                C.struct_ft
	Holog_cmd         C.struct_holog_cmd
	Ifp               C.struct_ifp
	In2net_cmd        C.struct_in2net_cmd
	In2net_mon        C.struct_in2net_mon
	K3fm_cmd          C.struct_k3fm_cmd
	K3fm_mon          C.struct_k3fm_mon
	K4label_cmd       C.struct_k4label_cmd
	K4pcalports_cmd   C.struct_k4pcalports_cmd
	K4pcalports_mon   C.struct_k4pcalports_mon
	K4rec_check       C.struct_k4rec_check
	K4rec_mode_cmd    C.struct_k4rec_mode_cmd
	K4rec_mode_mon    C.struct_k4rec_mode_mon
	K4recpatch_cmd    C.struct_k4recpatch_cmd
	K4st_cmd          C.struct_k4st_cmd
	K4vcbw_cmd        C.struct_k4vcbw_cmd
	K4vc_cmd          C.struct_k4vc_cmd
	K4vcif_cmd        C.struct_k4vcif_cmd
	K4vclo_cmd        C.struct_k4vclo_cmd
	K4vclo_mon        C.struct_k4vclo_mon
	K4vc_mon          C.struct_k4vc_mon
	Lo_cmd            C.struct_lo_cmd
	M5state           C.struct_m5state
	M5time            C.struct_m5time
	Mcb_cmd           C.struct_mcb_cmd
	Mcb_mon           C.struct_mcb_mon
	Mk5b_mode_cmd     C.struct_mk5b_mode_cmd
	Mk5b_mode_mon     C.struct_mk5b_mode_mon
	Monit5_ping       C.struct_monit5_ping
	Mux               C.struct_mux
	Onoff_cmd         C.struct_onoff_cmd
	Onoff_devices     C.struct_onoff_devices
	Out               C.struct_out
	Pcald_cmd         C.struct_pcald_cmd
	Pcalform_cmd      C.struct_pcalform_cmd
	Pcalports_cmd     C.struct_pcalports_cmd
	Pps_source_cmd    C.struct_pps_source_cmd
	Rclcn_req_buf     C.struct_rclcn_req_buf
	Rclcn_res_buf     C.struct_rclcn_res_buf
	Rec_mode_cmd      C.struct_rec_mode_cmd
	Regs              C.struct_regs
	Req_buf           C.struct_req_buf
	Req_rec           C.struct_req_rec
	Res_buf           C.struct_res_buf
	Res_rec           C.struct_res_rec
	Rtime_mon         C.struct_rtime_mon
	Rvac_cmd          C.struct_rvac_cmd
	Rvac_mon          C.struct_rvac_mon
	Rxgain_ds         C.struct_rxgain_ds
	S2bbc_data        C.struct_s2bbc_data
	S2das_check       C.struct_s2das_check
	S2label_cmd       C.struct_s2label_cmd
	S2_out            C.struct_s2_out
	S2rec_check       C.struct_s2rec_check
	S2st_cmd          C.struct_s2st_cmd
	Satellite_cmd     C.struct_satellite_cmd
	Satellite_ephem   C.struct_satellite_ephem
	Satoff_cmd        C.struct_satoff_cmd
	Scan_check_mon    C.struct_scan_check_mon
	Scan_name_cmd     C.struct_scan_name_cmd
	Servo             C.struct_servo
	Systracks_cmd     C.struct_systracks_cmd
	Tacd_shm          C.struct_tacd_shm
	Tape_cmd          C.struct_tape_cmd
	Tape_mon          C.struct_tape_mon
	Tle_cmd           C.struct_tle_cmd
	Tpicd_cmd         C.struct_tpicd_cmd
	User_device_cmd   C.struct_user_device_cmd
	User_info_cmd     C.struct_user_info_cmd
	User_info_parse   C.struct_user_info_parse
	Venable_cmd       C.struct_venable_cmd
	Vform_cmd         C.struct_vform_cmd
	Vform_mon         C.struct_vform_mon
	Vrepro_cmd        C.struct_vrepro_cmd
	Vsi4              C.struct_vsi4
	Vsi4_cmd          C.struct_vsi4_cmd
	Vsi4_mon          C.struct_vsi4_mon
	Vsn_mon           C.struct_vsn_mon
	Vst_cmd           C.struct_vst_cmd
	Wvolt_cmd         C.struct_wvolt_cmd
)
