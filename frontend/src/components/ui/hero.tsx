"use client";

import { Button, Link, Snippet } from "@nextui-org/react";
import { ChevronIcon } from "./icons";
import { TypeAnimation } from "react-type-animation";
import { useRouter } from "next/navigation";

export default function Hero() {
  const route = useRouter();
  return (
    <div className="text-center">
      <h1 className="text-gradient my-6 text-4xl md:text-7xl">
        <span className="bg-gradient-to-r from-teal-300 via-blue-500 to-blue-800 bg-clip-text text-transparent">
          &nbsp;Protect&nbsp;
        </span>
        your&nbsp;
        <TypeAnimation
          sequence={[
            2000, // Waits 2s
            "OpenAI credits", // Types 'One'
            2000, // Waits 1s
            "next.js routes", // Deletes 'One' and types 'Two'
            2000, // Waits 2s
            "API", // Types 'Three' without deleting 'Two'
            2000, // Waits 2s
          ]}
          wrapper="span"
          cursor={true}
          repeat={Infinity}
        />
      </h1>
      <Snippet className="mr-4">npm i @stratus-dev/sdk</Snippet>
      <Button
        as={Link}
        color="primary"
        variant="ghost"
        size="lg"
        href="/sign-up"
        style={{ position: "relative", zIndex: 1 }}
      >
        Get started
        <ChevronIcon />
      </Button>

      {/* fuzzy color in bg */}
      <div className="relative isolate px-6 pt-14 lg:px-8">
        <div
          className="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
          aria-hidden="true"
        >
          <div
            className="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-teal-300 to-blue-800 opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
            style={{
              clipPath:
                "polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)",
            }}
          />
        </div>
        <div className="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
          <div className="hidden sm:mb-8 sm:flex sm:justify-center">
            <div className="relative rounded-full px-3 py-1 text-sm leading-6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20">
              Announcing our closed beta.{" "}
              <a href="#" className="font-semibold text-blue-600">
                <span className="absolute inset-0" aria-hidden="true" />
                Read more <span aria-hidden="true">&rarr;</span>
              </a>
            </div>
          </div>
        </div>
        <div
          className="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
          aria-hidden="true"
        >
          <div
            className="relative left-[calc(50%+3rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 bg-gradient-to-tr from-teal-300 to-blue-800  opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
            style={{
              clipPath:
                "polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)",
            }}
          />
        </div>
      </div>
    </div>
  );
}
