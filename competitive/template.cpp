#include<bits/stdc++.h>

using namespace std;

#define FastIO           ios_base::sync_with_stdio(0)

#define Read(f)          freopen(f, "r", stdin)
#define Write(f)         freopen(f, "w", stdout)

#define Mem(a, b)        memset(a, b, sizeof a)
#define MemCp(a, b)      memcopy(a, b, sizeof a)
#define RADIANS(x)       (((1.0 * x * PI) / 180.0))
#define DEGREES(x)       (((x * 180.0) / (1.0 * pi)))
#define sq(x)            ((x)*(x))
#define debug(x)         cerr<<__LINE__<<" "<<#x<<" "<<x<<endl;

#define eps              1e-8
#define Pi               (2*acos(0.0))
#define INF              0x7f7f7f7f

#define in               scanf
#define out              printf
#define pb               push_back
#define popb             pop_back
#define Long             long long
#define ULong            unsigned long long
#define xx               first
#define yy               second
#define mp               make_pair

#define Mx               10000
#define Mod              10000

long dx[] = {-1, -1, -1,  0, 0,  1, 1, 1};
long dy[] = {-1,  0,  1, -1, 1, -1, 0, 1};

template < class T > T Abs(T a)
{
    return (a < 0 ? -a : a);
}

template < class T > T gcd(T a, T b)
{
    return (b != 0 ? gcd(b, a%b) : a);
}

template < class T > T lcm(T a, T b)
{
    return ((a*b)/gcd(a, b));
}

template < class T > T bigMod(T b, T p, T m)
{
     if(p==0)
          return 1;
     else if( p & 1 )
          return ( b*bigMod(b, p-1, m))%m;
     else
     {
          T temp = bigMod(b, p>>1, m);
          return (sq(temp))%m;
     }
}

template < class T > T modInverse(T b, T m)
{
     return bigMod(b, m-2, m);
}

template < class T > T power(T b, T p)
{
     T res=1, x=b;
     while(p)
     {
          if(p&1) res*=x;
          x*=x;
          p=p>>1;
     }
     return res;
}

template<typename T> inline bool isOn(T mask,T pos)
{
    return mask&(1<<pos);
}

template<typename T> inline int Off(T mask,T pos)
{
    return mask^(1<<pos);
}

template<typename T> inline int On(T mask,T pos)
{
    return mask|(1<<pos);
}

bool stringMod(string a, long long n)
{
     long long num=0;
     for(long i=0; i<a.size(); i++)
          num=((num*10) + (a[i]-'0'))%n;

     if(num==0)
          return 1;
     else
          return 0;
}


bool prime[Mx+2];
vector <long long> primes;
void seive()
{
     for(long long i=4; i<=Mx; i+=2)
          prime[i]=1;

     for(long long i=3; i*i<=Mx; i+=2)
     {
          if(prime[i]==0)
          {
               for(long long j=2*i; j<=Mx; j+=i)
                    prime[j]=1;
          }
     }

     primes.pb(2);
     for(long long i=3; i<=Mx; i+=2)
          if(prime[i]==0)
               primes.pb(i);
}


void solve()
{

}

int main()
{
    FastIO;
    solve();
}