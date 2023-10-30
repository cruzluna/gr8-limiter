import Hero from "@/components/ui/hero";
import StratusNavbar from "@/components/ui/navbar";

export default function Home() {
  return (
    <div>
      <section className="overflow-hidden pb-[16.4rem] md:pb-[25.6rem] text-white">
        <div className="pt-[6.4rem]">
          <Hero />
        </div>
      </section>
    </div>
  );
}
