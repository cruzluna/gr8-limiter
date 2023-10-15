"use client";
import Hero from "@/components/ui/hero";
import StratusNavbar from "@/components/ui/navbar";
import { Button } from "@nextui-org/react";
import { useUser } from "@clerk/nextjs";
import { v4 as uuidv4 } from "uuid";

export default function Home() {
  const { user, isLoaded } = useUser(); // get clerk user for clerkId

  // POC
  const handleSubmit = async () => {
    console.log("in here");

    if (isLoaded) {
      try {
        await fetch("/api/apikey", {
          method: "POST",
          body: JSON.stringify({
            api_key: uuidv4(),

            user_id: user?.id,
          }),
          headers: {
            "Content-Type": "application/json",
          },
        });
      } catch (error) {
        console.error(error);
      }
    }
  };
  return (
    <div>
      <section>
        <StratusNavbar />
      </section>
      <section className="overflow-hidden pb-[16.4rem] md:pb-[25.6rem] text-white">
        <div className="pt-[6.4rem]">
          <Hero />
        </div>
      </section>
      {/*POC*/}
      <Button onClick={handleSubmit}>test</Button>
    </div>
  );
}
