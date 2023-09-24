"use client";

import { Button, Link } from "@nextui-org/react";
import { ChevronIcon } from "./icons";

export default function Hero() {
  return (
    <div className="text-center">
      <h1 className="text-gradient my-6 text-4xl md:text-7xl">
        make
        <span className="bg-gradient-to-r from-teal-300 via-blue-500 to-blue-800 bg-clip-text text-transparent">
          &nbsp;rate&nbsp;
        </span>
        limiting easy
      </h1>
      <Button
        as={Link}
        color="primary"
        variant="ghost"
        size="lg"
        href="/sign-up"
      >
        Get started
        <ChevronIcon />
      </Button>
    </div>
  );
}
