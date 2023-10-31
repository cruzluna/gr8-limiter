import { neon } from "@neondatabase/serverless";

if (!process.env.DB_URL) {
  throw new Error("DB_URL not set");
}

export const sql = neon(process.env.DB_URL) || "";
