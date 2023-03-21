#ifndef RTKCMN_H
#define RTKCMN_H
#include <stdarg.h>
#include <ctype.h>
#include <errno.h>
#ifndef WIN32
#include <dirent.h>
#include <time.h>
#include <sys/time.h>
#include <sys/stat.h>
#include <sys/types.h>
#endif
#include "rtcm.h"

#ifdef __cplusplus
extern "C" {
#endif

/* constants -----------------------------------------------------------------*/
#define POLYCRC32   0xEDB88320u /* CRC32 polynomial */
#define POLYCRC24Q  0x1864CFBu  /* CRC24Q polynomial */

#define SQR(x)      ((x)*(x))
#define MAX_VAR_EPH SQR(300.0)  /* max variance eph to reject satellite (m^2) */

typedef void fatalfunc_t(const char *); /* fatal callback function type */

EXPORT void add_fatal(fatalfunc_t *func);

/* satellites, systems, codes functions --------------------------------------*/
EXPORT int  satno   (int sys, int prn);
EXPORT int  satsys  (int sat, int *prn);
EXPORT int  satid2no(const char *id);
EXPORT void satno2id(int sat, char *id);
EXPORT uint8_t obs2code(const char *obs);
EXPORT char *code2obs(uint8_t code);
EXPORT double code2freq(int sys, uint8_t code, int fcn);
EXPORT double sat2freq(int sat, uint8_t code, const nav_t *nav);
EXPORT int  code2idx(int sys, uint8_t code);
EXPORT void setcodepri(int sys, int idx, const char *pri);
EXPORT int  getcodepri(int sys, uint8_t code, const char *opt);

/* receiver raw data functions -----------------------------------------------*/
EXPORT uint32_t getbitu(const uint8_t *buff, int pos, int len);
EXPORT int32_t  getbits(const uint8_t *buff, int pos, int len);
EXPORT void setbitu(uint8_t *buff, int pos, int len, uint32_t data);
EXPORT void setbits(uint8_t *buff, int pos, int len, int32_t  data);
EXPORT uint32_t rtk_crc24q(const uint8_t *buff, int len);

/* debug trace functions -----------------------------------------------------*/
EXPORT void trace    (int level, const char *format, ...);

/* coordinates transformation ------------------------------------------------*/
EXPORT void ecef2pos(const double *r, double *pos);
EXPORT void pos2ecef(const double *pos, double *r);
EXPORT void ecef2enu(const double *pos, const double *r, double *e);
EXPORT void enu2ecef(const double *pos, const double *e, double *r);
EXPORT void covenu  (const double *pos, const double *P, double *Q);
EXPORT void covecef (const double *pos, const double *Q, double *P);
EXPORT void xyz2enu (const double *pos, double *E);
EXPORT void eci2ecef(gtime_t tutc, const double *erpv, double *U, double *gmst);
EXPORT void deg2dms (double deg, double *dms, int ndec);
EXPORT double dms2deg(const double *dms);

EXPORT gtime_t timeadd  (gtime_t t, double sec);
EXPORT double  timediff (gtime_t t1, gtime_t t2);
EXPORT gtime_t gpst2utc (gtime_t t);
EXPORT gtime_t utc2gpst (gtime_t t);
EXPORT gtime_t gpst2bdt (gtime_t t);
EXPORT gtime_t bdt2gpst (gtime_t t);

EXPORT gtime_t timeget  (void);
EXPORT void    timeset  (gtime_t t);
EXPORT void    timereset(void);
EXPORT double  time2doy (gtime_t t);
EXPORT double  utc2gmst (gtime_t t, double ut1_utc);

EXPORT int adjgpsweek(int week);
EXPORT uint32_t tickget(void);
EXPORT void sleepms(int ms);

/* time and string functions -------------------------------------------------*/
EXPORT double  str2num(const char *s, int i, int n);
EXPORT int     str2time(const char *s, int i, int n, gtime_t *t);
EXPORT void    time2str(gtime_t t, char *str, int n);
EXPORT gtime_t epoch2time(const double *ep);
EXPORT void    time2epoch(gtime_t t, double *ep);
EXPORT gtime_t gpst2time(int week, double sec);
EXPORT double  time2gpst(gtime_t t, int *week);
EXPORT gtime_t gst2time(int week, double sec);
EXPORT double  time2gst(gtime_t t, int *week);
EXPORT gtime_t bdt2time(int week, double sec);
EXPORT double  time2bdt(gtime_t t, int *week);
EXPORT char    *time_str(gtime_t t, int n);

/* matrix and vector functions -----------------------------------------------*/
EXPORT double *mat  (int n, int m);
EXPORT int    *imat (int n, int m);
EXPORT double *zeros(int n, int m);
EXPORT double *eye  (int n);
EXPORT double dot (const double *a, const double *b, int n);
EXPORT double norm(const double *a, int n);
EXPORT void cross3(const double *a, const double *b, double *c);
EXPORT int  normv3(const double *a, double *b);
EXPORT void matcpy(double *A, const double *B, int n, int m);
EXPORT void matmul(const char *tr, int n, int k, int m, double alpha,
                   const double *A, const double *B, double beta, double *C);
EXPORT int  matinv(double *A, int n);
EXPORT int  solve (const char *tr, const double *A, const double *Y, int n,
                   int m, double *X);
EXPORT int  lsq   (const double *A, const double *y, int n, int m, double *x,
                   double *Q);
EXPORT int  filter(double *x, double *P, const double *H, const double *v,
                   const double *R, int n, int m);
EXPORT int  smoother(const double *xf, const double *Qf, const double *xb,
                     const double *Qb, int n, double *xs, double *Qs);
EXPORT void matprint (const double *A, int n, int m, int p, int q);
EXPORT void matfprint(const double *A, int n, int m, int p, int q, FILE *fp);


#ifdef __cplusplus
}
#endif

#endif /* RTKCMN_H */
