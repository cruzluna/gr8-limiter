# Gr8 Limiter

<p align="center">
  Throw in your next.js project for rate limiting...
</p>

<div align="center">
  <a href="https://gr8-limiter.vercel.app/">🔗Home</a> | <a href="https://stratus-docs-production.up.railway.app"/>📝Docs</a> | <a href="https://github.com/cruzluna/gr8-limiter/blob/main/backend/internal/docs/general-flow.png">🛠️Design</a>
</div>


A drop in rate limiter designed for web frameworks like next.js. Focus on shipping code and Gr8 limiter will 
handle rate limiting with limited to no set up.



# How to use?
All the code you need: 

`route.ts`
```typescript

const client = new Client({
  // this will be cleaned up soon...
  apiKey: process.env.STRATUS_TOKEN!,
  apiURL: "https://gr8-limit-docker.onrender.com/api/v1/ratelimit",
});

export async function POST(req: Next Request){
// Inside a route method
  try {
    const rateLimited = await client.rateLimit();
    if (rateLimited) {
      return NextResponse.json(
        { error: "Rate limited." },
        { status: 429 }
      );
    }
    // Rest of ur code...
```

### Stack: 
<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.typescriptlang.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/typescript/typescript-original.svg" alt="typescript" width="40" height="40"/> </a> <a href="https://nextjs.org/" target="_blank" rel="noreferrer"> <img src="https://cdn.worldvectorlogo.com/logos/nextjs-2.svg" alt="nextjs" width="40" height="40"/> </a> <a href="https://www.postgresql.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-original-wordmark.svg" alt="postgresql" width="40" height="40"/> </a> <a href="https://redis.io" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/redis/redis-original-wordmark.svg" alt="redis" width="40" height="40"/> </a> <a href="https://tailwindcss.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/tailwindcss/tailwindcss-icon.svg" alt="tailwind" width="40" height="40"/> </a>  </p>
