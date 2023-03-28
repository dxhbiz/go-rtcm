package rtcm

/*
#cgo CFLAGS: -I./ -DENAGLO -DENAGAL -DENAQZS -DENACMP -DENAIRN

#include "rtkcmn.h"
#include "rtcm.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	ErrIncomplete = errors.New("incomplete data")
	ErrCrc        = errors.New("crc error")
	ErrInvalid    = errors.New("invalid data")
)

// GtimeT time struct
type GtimeT struct {
	Time int64   // time (s)
	Sec  float64 // fraction of second under 1 s
}

// ObsdT observation data record
type ObsdT struct {
	Time GtimeT    // receiver sampling time (GPST)
	Sat  uint8     // satellite number
	Rcv  uint8     // receiver number
	SNR  []uint16  // signal strength (0.001 dBHz)
	LLI  []uint8   // loss of lock indicator
	Code []uint8   // code indicator (CODE_???)
	L    []float64 // observation data carrier-phase (cycle)
	P    []float64 // observation data pseudorange (m)
	D    []float32 // observation data doppler frequency (Hz)
}

// ObsT observation data
type ObsT struct {
	N    int     // number of obervation data/allocated
	Nmax int     // max number of obervation data/allocated
	Data []ObsdT // observation data records
}

// EphT GPS/QZS/GAL broadcast ephemeris type
type EphT struct {
	Sat  int // satellite number
	Iode int // IODE
	Iodc int // IODC
	Sva  int // SV accuracy (URA index)
	Svh  int // SV health (0:ok)
	Week int // GPS/QZS: gps week, GAL: galileo week
	Code int // GPS/QZS: code on L2
	// GAL: data source defined as rinex 3.03
	// BDS: data source (0:unknown,1:B1I,2:B1Q,3:B2I,4:B2Q,5:B3I,6:B3Q)
	Flag int // GPS/QZS: L2 P data flag
	// BDS: nav type (0:unknown,1:IGSO/MEO,2:GEO)
	Toe  GtimeT // Toe
	Toc  GtimeT // Toc
	Ttr  GtimeT // T_trans
	A    float64
	E    float64
	Io   float64
	OMG0 float64
	Omg  float64
	M0   float64
	Deln float64
	OMGd float64
	Idot float64
	Crc  float64
	Crs  float64
	Cuc  float64
	Cus  float64
	Cic  float64
	Cis  float64
	Toes float64 // Toe (s) in week
	Fit  float64 // fit interval (h)
	F0   float64 // SV clock parameters (af0,af1,af2)
	F1   float64
	F2   float64
	Tgd  []float64 // group delay parameters
	/* GPS/QZS:tgd[0]=TGD */
	/* GAL:tgd[0]=BGD_E1E5a,tgd[1]=BGD_E1E5b */
	/* CMP:tgd[0]=TGD_B1I ,tgd[1]=TGD_B2I/B2b,tgd[2]=TGD_B1Cp */
	/*     tgd[3]=TGD_B2ap,tgd[4]=ISC_B1Cd   ,tgd[5]=ISC_B2ad */
	Adot float64 // adot for CNAV
	NDot float64 // ndot for CNAV
}

// GephT GLONASS broadcast ephemeris type
type GephT struct {
	Sat   int        // satellite number
	Iode  int        // IODE (0-6 bit of tb field)
	Frq   int        // satellite frequency number
	Svh   int        // satellite health
	Sva   int        // satellite accuracy
	Age   int        // satellite age of operation
	Toe   GtimeT     // epoch of epherides (gpst)
	Tof   GtimeT     // message frame time (gpst)
	Pos   [3]float64 // satellite position (ecef) (m)
	Vel   [3]float64 // satellite velocity (ecef) (m/s)
	Acc   [3]float64 // satellite acceleration (ecef) (m/s^2)
	Taun  float64    // SV clock bias (s)
	Gamn  float64    // relative freq bias
	Dtaun float64    // delay between L1 and L2 (s)
}

// SephT SBAS ephemeris type
type SephT struct {
	Sat int        // satellite number
	T0  GtimeT     // reference epoch time (GPST)
	Tof GtimeT     // time of message frame (GPST)
	Sva int        // SV accuracy (URA index)
	Svh int        // SV health (0:ok)
	Pos [3]float64 // satellite position (m) (ecef)
	Vel [3]float64 // satellite velocity (m/s) (ecef)
	Acc [3]float64 // satellite acceleration (m/s^2) (ecef)
	Af0 float64    // satellite clock-offset (s)
	Af1 float64    // satellite drift (s/s)
}

// PephT precise ephemeris type
type PephT struct {
	Time  GtimeT       // time (GPST)
	Index int          // ephemeris index for multiple files
	Pos   [][4]float64 // satellite position/clock (ecef) (m|s)
	Std   [][4]float32 // satellite position/clock std (m|s)
	Vel   [][4]float64 // satellite velocity/clk-rate (m/s|s/s)
	Vst   [][4]float32 // satellite velocity/clk-rate std (m/s|s/s)
	Cov   [][3]float32 // satellite position covariance (m^2)
	Vcos  [][3]float32 // satellite velocity covariance (m^2)
}

// PclkT precise clock type
type PclkT struct {
	Time  GtimeT       // time (GPST)
	Index int          // clock index for multiple files
	Clk   [][1]float64 // satellite clock (s)
	Std   [][1]float32 // satellite clock std (s)
}

// AlmT almanac type
type AlmT struct {
	Sat    int    // satellite number
	Svh    int    // sv health (0:ok)
	Svconf int    // as and sv config
	Week   int    // GPS/QZS: gps week, GAL: galileo week
	Toa    GtimeT // Toa
	// SV orbit parameters
	A    float64
	E    float64
	I0   float64
	OMG0 float64
	Omg  float64
	M0   float64
	OMGd float64
	Toas float64 // Toa (s) in week
	F0   float64 // SV clock parameters (af0,af1)
	F1   float64
}

// TecT TEC grid type
type TecT struct {
	Time  GtimeT     // epoch time (GPST)
	Ndata [3]int     // TEC grid data size {nlat,nlon,nhgt}
	Rb    float64    // earth radius (km)
	Lats  [3]float64 // latitude start/interval (deg)
	Lons  [3]float64 // longitude start/interval (deg)
	Hgts  [3]float64 // heights start/interval (km)
	Data  []float64  // TEC grid data (tecu)
	Rms   []float32  // RMS values (tecu)
}

// ErpData earth rotation parameter data type
type ErpData struct {
	Mjd    float64 // mjd (days)
	Xp     float64 // x pole offset (rad)
	Yp     float64 // y pole offset (rad)
	Xpr    float64 // x pole offset rate (rad/day)
	Ypr    float64 // y pole offset rate (rad/day)
	Ut1Utc float64 // ut1-utc (s)
	Lod    float64 // length of day (s/day)
}

// ErpT earth rotation parameter type
type ErpT struct {
	N    int       // number of data
	Nmax int       // max number of data
	Data []ErpData // earth rotation parameter data
}

// PcvT antenna parameter type
type PcvT struct {
	Sat  int          // satellite number (0:receiver)
	Type []int8       // antenna type
	Code []int8       // serial number or satellite code
	Ts   GtimeT       // valid time start
	Te   GtimeT       // valid time end
	Off  [][3]float64 // phase center offset e/n/u or x/y/z (m)
	Var  [][3]float64 // phase center variation (m), el=90,85,...,0 or nadir=0,1,2,3,... (deg)
}

// SbsfcorrT SBAS fast correction type
type SbsfcorrT struct {
	T0   GtimeT  // time of applicability (TOF)
	Prc  float64 // pseudorange correction (PRC) (m)
	Rrc  float64 // range-rate correction (RRC) (m/s)
	Dt   float64 // range-rate correction delta-time (s)
	Iodf int     // IODF (issue of date fast corr)
	Udre int16   // UDRE+1
	Ai   int16   // degradation factor indicator
}

// SbslcorrT SBAS long term satellite error correction type
type SbslcorrT struct {
	T0   GtimeT     // correction time
	Iode int        // IODE (issue of date ephemeris)
	Dpos [3]float64 // delta position (m) (ecef)
	Dvel [3]float64 // delta velocity (m/s) (ecef)
	Daf0 float64    // delta clock-offset/drift (s,s/s)
	Daf1 float64    // delta clock-offset/drift (s,s/s)
}

// SbssatpT SBAS satellite correction type
type SbssatpT struct {
	Sat   int       // satellite number
	Fcorr SbsfcorrT // fast correction
	Lcorr SbslcorrT // long term correction
}

// SbssatT SBAS satellite corrections type
type SbssatT struct {
	Iodp int        // IODP (issue of date mask)
	Nsat int        // number of satellites
	Tlat int        // system latency (s)
	Sat  []SbssatpT // satellite correction
}

// SbsigpT SBAS ionospheric correction type
type SbsigpT struct {
	T0    GtimeT  // correction time
	Lat   int16   // latitude (deg)
	Lon   int16   // longitude (deg)
	Give  int16   // GIVI+1
	Delay float32 // vertical delay estimate (m)
}

// SbsionT SBAS ionospheric corrections type
type SbsionT struct {
	Iodi int       // IODI (issue of date ionos corr)
	Nigp int       // number of igps
	Igp  []SbsigpT // ionospheric correction
}

// DgpsT DGPS/GNSS correction type
type DgpsT struct {
	T0   GtimeT  // correction time
	Prc  float64 // pseudorange correction (PRC) (m)
	Rrc  float64 // range rate correction (RRC) (m/s)
	Iod  int     // issue of data (IOD)
	Udre float32 // UDRE
}

// NavT navigation data type
type NavT struct {
	N      int             // number of broadcast ephemeris
	Nmax   int             // max number of broadcast ephemeris
	Ng     int             // number of glonass ephemeris
	NgMax  int             // max number of glonass ephemeris
	Ns     int             // number of sbas ephemeris
	NsMax  int             // max number of sbas ephemeris
	Ne     int             // number of precise ephemeris
	NeMax  int             // max number of precise ephemeris
	Nc     int             // number of precise clock
	NcMax  int             // max number of precise clock
	Na     int             // number of almanac data
	NaMax  int             // max number of almanac data
	Nt     int             // number of tec grid data
	NtMax  int             // max number of tec grid data
	Eph    []EphT          // GPS/QZS/GAL/BDS/IRN ephemeris
	Geph   []GephT         // GLONASS ephemeris
	Seph   []SephT         // SBAS ephemeris
	Peph   []PephT         // precise ephemeris
	Pclk   []PclkT         // precise clock
	Alm    []AlmT          // almanac data
	Tec    []TecT          // tec grid data
	Erp    []ErpT          // earth rotation parameters
	UtcGps [8]float64      // GPS delta-UTC parameters {A0,A1,Tot,WNt,dt_LS,WN_LSF,DN,dt_LSF}
	UtcGlo [8]float64      // GLONASS UTC time parameters {tau_C,tau_GPS}
	UtcGal [8]float64      // Galileo UTC parameters
	UtcQzs [8]float64      // QZS UTC parameters
	UtcCmp [8]float64      // BeiDou UTC parameters
	UtcIrn [9]float64      // IRNSS UTC parameters {A0,A1,Tot,...,dt_LSF,A2}
	UtcSbs [4]float64      // SBAS UTC parameters
	IonGps [8]float64      // GPS iono model parameters {a0,a1,a2,a3,b0,b1,b2,b3}
	IonGal [4]float64      // Galileo iono model parameters {ai0,ai1,ai2,0}
	IonQzs [8]float64      // QZSS iono model parameters {a0,a1,a2,a3,b0,b1,b2,b3}
	IonCmp [8]float64      // BeiDou iono model parameters {a0,a1,a2,a3,b0,b1,b2,b3}
	IonIrn [8]float64      // IRNSS iono model parameters {a0,a1,a2,a3,b0,b1,b2,b3}
	GloFcn [32]int         // GLONASS FCN + 8
	Cbias  [][3]float64    // satellite DCB (0:P1-P2,1:P1-C1,2:P2-C2) (m)
	Rbias  [][2][3]float64 // receiver DCB (0:P1-P2,1:P1-C1,2:P2-C2) (m)
	Pcvs   []PcvT          // satellite antenna pcv
	Sbssat SbssatT         // SBAS satellite corrections
	Sbsion []SbsionT       // SBAS ionosphere corrections
	Dgps   []DgpsT         // DGPS corrections
	SSR    []SSRT          // SSR corrections
}

// StaT station parameter type
type StaT struct {
	Name       string     // marker name
	Marker     string     // marker number
	Antdes     string     // antenna descriptor
	Antsno     string     // antenna serial number
	Rectype    string     // receiver type descriptor
	Recver     string     // receiver firmware version
	Recsno     string     // receiver serial number
	Antsetup   int        // antenna setup id
	Itrf       int        // ITRF realization year
	Deltype    int        //  antenna delta type (0:enu,1:xyz)
	Pos        [3]float64 // station position (ecef) (m)
	Del        [3]float64 // antenna position delta (e/n/u or x/y/z) (m)
	Hgt        float64    // antenna height (m)
	GloCpAlign int        // GLONASS code-phase alignment (0:no,1:yes)
	GloCpBias  [4]float64 // GLONASS code-phase biases {1C,1P,2C,2P} (m)
}

type SSRT struct {
	T0      [6]GtimeT  // epoch time (GPST) {eph,clk,hrclk,ura,bias,pbias}
	Udi     [6]float64 // SSR update interval (s)
	Iod     [6]int     // iod ssr {eph,clk,hrclk,ura,bias,pbias}
	Iode    int        // issue of data
	Iodcrc  int        // issue of data crc for beidou/sbas
	Ura     int        // URA indicator
	Refd    int        // sat ref datum (0:ITRF,1:regional)
	Deph    [3]float64 // delta orbit {radial,along,cross} (m)
	Ddeph   [3]float64 // dot delta orbit {radial,along,cross} (m/s)
	Dclck   [3]float64 // delta clock {c0,c1,c2} (m,m/s,m/s^2)
	Hrclk   float64    // high-rate clock corection (m)
	Cbias   []float32  // code biases (m)
	Pbias   []float64  // phase biases (m)
	Stdpb   []float32  // std-dev of phase biases (m)
	YawAng  float64    // yaw angle (deg)
	YawRate float64    // yaw rate (deg/s)
	Update  uint8      // update flag (0:no update,1:update)
}

// Rtcm RTCM control struct type
type Rtcm struct {
	Type    int     // message type
	Crc     int     // CRC value, 0 means no CRC error, 1 means has CRC error
	StaId   int     // station id
	Time    GtimeT  // message time
	Obs     ObsT    // observation data
	Nav     NavT    // satellite ephemerides
	Sta     StaT    // station parameters
	Dgps    []DgpsT // output of dgps corrections
	SSR     []SSRT  // output of ssr corrections
	ObsFlag int     // obs data complete flag (1:ok,0:not complete)
	EphSat  int     // input ephemeris satellite number
	EphSet  int     // input ephemeris set (0-1)
}

// Cdecode decode the rtcm data and return the C.rtcm_t
func Cdecode(data []byte) (rtcm C.rtcm_t, err error) {
	C.init_rtcm(&rtcm)
	for _, d := range data {
		C.input_rtcm3(&rtcm, C.uint8_t(d))
	}
	C.free_rtcm(&rtcm)
	return
}

// convertObs convert C.obs_t to go ObsT
func convertObs(cobs C.obs_t) ObsT {
	obs := ObsT{
		N:    int(cobs.n),
		Nmax: int(cobs.nmax),
	}

	if obs.N > 0 {
		var dataArr *C.obsd_t = cobs.data
		dataSlice := unsafe.Slice(dataArr, obs.N)
		eleLen := int(C.NFREQ) + int(C.NEXOBS)

		for i := 0; i < obs.N; i++ {
			cobsd := dataSlice[i]

			time := GtimeT{
				Time: int64(cobsd.time.time),
				Sec:  float64(cobsd.time.sec),
			}

			snr := make([]uint16, eleLen)
			lli := make([]uint8, eleLen)
			code := make([]uint8, eleLen)
			l := make([]float64, eleLen)
			p := make([]float64, eleLen)
			d := make([]float32, eleLen)
			for j := 0; j < eleLen; j++ {
				snr[j] = uint16(cobsd.SNR[j])
				lli[j] = uint8(cobsd.LLI[j])
				code[j] = uint8(cobsd.code[j])
				l[j] = float64(cobsd.L[j])
				p[j] = float64(cobsd.P[j])
				d[j] = float32(cobsd.D[j])
			}

			obsd := ObsdT{
				Time: time,
				Sat:  uint8(cobsd.sat),
				Rcv:  uint8(cobsd.rcv),
				SNR:  snr,
				LLI:  lli,
				Code: code,
				L:    l,
				P:    p,
				D:    d,
			}

			obs.Data = append(obs.Data, obsd)
		}
	}

	return obs
}

// convertNav convert C.nav_t to go NavT
func convertNav(cnav C.nav_t) NavT {
	nav := NavT{
		N:     int(cnav.n),
		Nmax:  int(cnav.nmax),
		Ng:    int(cnav.ng),
		NgMax: int(cnav.ngmax),
		Ns:    int(cnav.ns),
		NsMax: int(cnav.nsmax),
		Ne:    int(cnav.ne),
		NeMax: int(cnav.nemax),
		Nc:    int(cnav.nc),
		Na:    int(cnav.na),
		NaMax: int(cnav.namax),
		Nt:    int(cnav.nt),
		NtMax: int(cnav.ntmax),
	}

	// copy GPS/QZS/GAL/BDS/IRN ephemeris
	if nav.N > 0 {
		var dataArr *C.eph_t = cnav.eph
		dataSlice := unsafe.Slice(dataArr, nav.N)

		for i := 0; i < nav.N; i++ {
			ceph := dataSlice[i]

			if ceph.sat <= 0 {
				continue
			}

			eph := EphT{
				Sat:  int(ceph.sat),
				Iode: int(ceph.iode),
				Iodc: int(ceph.iodc),
				Sva:  int(ceph.sva),
				Svh:  int(ceph.svh),
				Week: int(ceph.week),
				Code: int(ceph.code),
				Flag: int(ceph.flag),
				Toe: GtimeT{
					Time: int64(ceph.toe.time),
					Sec:  float64(ceph.toe.sec),
				},
				Toc: GtimeT{
					Time: int64(ceph.toc.time),
					Sec:  float64(ceph.toc.sec),
				},
				Ttr: GtimeT{
					Time: int64(ceph.ttr.time),
					Sec:  float64(ceph.ttr.sec),
				},
				A:    float64(ceph.A),
				E:    float64(ceph.e),
				Io:   float64(ceph.i0),
				OMG0: float64(ceph.OMG0),
				Omg:  float64(ceph.omg),
				M0:   float64(ceph.M0),
				Deln: float64(ceph.deln),
				OMGd: float64(ceph.OMGd),
				Idot: float64(ceph.idot),
				Crc:  float64(ceph.crc),
				Crs:  float64(ceph.crs),
				Cuc:  float64(ceph.cuc),
				Cus:  float64(ceph.cus),
				Cic:  float64(ceph.cic),
				Cis:  float64(ceph.cis),
				Toes: float64(ceph.toes),
				Fit:  float64(ceph.fit),
				F0:   float64(ceph.f0),
				F1:   float64(ceph.f1),
				F2:   float64(ceph.f2),
				Adot: float64(ceph.Adot),
				NDot: float64(ceph.ndot),
			}

			tgd := make([]float64, 6)
			for j := 0; j < 6; j++ {
				tgd[j] = float64(ceph.tgd[j])
			}
			eph.Tgd = tgd

			nav.Eph = append(nav.Eph, eph)
		}
	}

	// copy GLONASS ephemeris
	if nav.Ng > 0 {
		var dataArr *C.geph_t = cnav.geph
		dataSlice := unsafe.Slice(dataArr, nav.Ng)

		for i := 0; i < nav.Ng; i++ {
			cgeph := dataSlice[i]

			if cgeph.sat <= 0 {
				continue
			}

			geph := GephT{
				Sat:  int(cgeph.sat),
				Iode: int(cgeph.iode),
				Frq:  int(cgeph.frq),
				Svh:  int(cgeph.svh),
				Sva:  int(cgeph.sva),
				Age:  int(cgeph.age),
				Toe: GtimeT{
					Time: int64(cgeph.toe.time),
					Sec:  float64(cgeph.toe.sec),
				},
				Tof: GtimeT{
					Time: int64(cgeph.tof.time),
					Sec:  float64(cgeph.tof.sec),
				},
				Taun:  float64(cgeph.taun),
				Gamn:  float64(cgeph.gamn),
				Dtaun: float64(cgeph.dtaun),
			}

			for j := 0; j < 3; j++ {
				geph.Pos[j] = float64(cgeph.pos[j])
				geph.Vel[j] = float64(cgeph.vel[j])
				geph.Acc[j] = float64(cgeph.acc[j])
			}

			nav.Geph = append(nav.Geph, geph)
		}
	}

	return nav
}

// convertSta convert C.sta_t to go StaT
func convertSta(csta C.sta_t) StaT {
	sta := StaT{
		Name:       C.GoString(&csta.name[0]),
		Marker:     C.GoString(&csta.marker[0]),
		Antdes:     C.GoString(&csta.antdes[0]),
		Antsno:     C.GoString(&csta.antsno[0]),
		Rectype:    C.GoString(&csta.rectype[0]),
		Recver:     C.GoString(&csta.recver[0]),
		Recsno:     C.GoString(&csta.recsno[0]),
		Antsetup:   int(csta.antsetup),
		Itrf:       int(csta.itrf),
		Deltype:    int(csta.deltype),
		Hgt:        float64(csta.hgt),
		GloCpAlign: int(csta.glo_cp_align),
	}

	for i := 0; i < 4; i++ {
		if i < 3 {
			sta.Pos[i] = float64(csta.pos[i])
			sta.Del[i] = float64(csta.del[i])
		}
		sta.GloCpBias[i] = float64(csta.glo_cp_bias[i])
	}

	return sta
}

// convertSSR convert [MAXSAT]C.ssr_t to go []SSRT
func convertSSR(cssr [C.MAXSAT]C.ssr_t) []SSRT {
	ssr := []SSRT{}

	for i := 0; i < int(C.MAXSAT); i++ {
		cssrt := cssr[i]

		if cssrt.update == 0 {
			continue
		}

		gssr := SSRT{
			Iode:    int(cssrt.iode),
			Iodcrc:  int(cssrt.iodcrc),
			Ura:     int(cssrt.ura),
			Refd:    int(cssrt.refd),
			Hrclk:   float64(cssrt.hrclk),
			YawAng:  float64(cssrt.yaw_ang),
			YawRate: float64(cssrt.yaw_rate),
			Update:  uint8(cssrt.update),
		}

		for j := 0; j < 6; j++ {
			gssr.T0[j] = GtimeT{
				Time: int64(cssrt.t0[j].time),
				Sec:  float64(cssrt.t0[j].sec),
			}

			gssr.Udi[j] = float64(cssrt.udi[j])
			gssr.Iod[j] = int(cssrt.iod[j])

			if j < 3 {
				gssr.Deph[j] = float64(cssrt.deph[j])
				gssr.Ddeph[j] = float64(cssrt.ddeph[j])
				gssr.Dclck[j] = float64(cssrt.dclk[j])
			}
		}

		maxCode := int(C.MAXCODE)
		gssr.Cbias = make([]float32, maxCode)
		gssr.Pbias = make([]float64, maxCode)
		gssr.Stdpb = make([]float32, maxCode)

		for j := 0; j < maxCode; j++ {
			gssr.Cbias[j] = float32(cssrt.cbias[j])
			gssr.Pbias[j] = float64(cssrt.pbias[j])
			gssr.Stdpb[j] = float32(cssrt.stdpb[j])
		}

		ssr = append(ssr, gssr)
	}

	return ssr
}

// Decode decode the rtcm data and return the Rtcm struct
func Decode(data []byte) (rtcm Rtcm, err error) {
	var crtcm C.rtcm_t
	C.init_rtcm(&crtcm)

	var status int
	for _, d := range data {
		status = int(C.input_rtcm3(&crtcm, C.uint8_t(d)))
	}

	rtcm.Type = int(crtcm.mtype)
	rtcm.Crc = int(crtcm.crc)
	rtcm.StaId = int(crtcm.staid)

	if status == -1 {
		err = ErrIncomplete
		return
	}

	if rtcm.Type == 0 && status == 0 {
		err = ErrInvalid
		return
	}

	if rtcm.Crc == 1 {
		err = ErrCrc
		return
	}

	rtcm.Time = GtimeT{
		Time: int64(crtcm.time.time),
		Sec:  float64(crtcm.time.sec),
	}

	rtcm.Obs = convertObs(crtcm.obs)
	rtcm.Nav = convertNav(crtcm.nav)
	rtcm.Sta = convertSta(crtcm.sta)
	rtcm.SSR = convertSSR(crtcm.ssr)

	rtcm.ObsFlag = int(crtcm.obsflag)
	rtcm.EphSat = int(crtcm.ephsat)
	rtcm.EphSet = int(crtcm.ephset)

	C.free_rtcm(&crtcm)
	return
}

// CheckCrc test whether rtcm3 data crc checksum is normal
func CheckCrc(data []byte) bool {
	dataLen := len(data)
	if dataLen < 6 {
		return false
	}

	checkLen := dataLen - 3
	cdata := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	checkVal := C.rtk_crc24q(cdata, C.int(checkLen))
	crcVal := C.getbitu(cdata, C.int(checkLen*8), 24)

	return checkVal == crcVal
}

// EcefToLla transform ecef to geodetic postion
func EcefToLla(ecef [3]float64) [3]float64 {
	var llh [3]float64
	cEcef := (*C.double)(unsafe.Pointer(&ecef[0]))
	cPos := (*C.double)(unsafe.Pointer(&llh[0]))
	C.ecef2lla(cEcef, cPos)
	return llh
}
