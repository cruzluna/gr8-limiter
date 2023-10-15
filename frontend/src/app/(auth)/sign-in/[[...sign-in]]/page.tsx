import { SignIn } from "@clerk/nextjs";

export default function Page() {
  return (
    <div className="flex items-center justify-center pt-8">
      <SignIn redirectUrl={"/dashboard"} />
    </div>
  );
}
