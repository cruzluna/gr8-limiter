"use client";

import { Card, CardHeader } from "@nextui-org/react";

export default function DashboardTiles() {
  return (
    <section>
      <div className="flex items-center justify-center pt-[1rem]">
        <div className="flex flex-row flex-grow-0 gap-x-4 w-full p-3">
          <Card className="w-1/3 py-4" isPressable isHoverable>
            <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
              <small className="text-default-500">12 Integrations</small>
              <h4 className="font-bold text-large">Manage Integrations</h4>
            </CardHeader>
          </Card>

          <Card className="w-1/3 py-4" isPressable isHoverable>
            <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
              <small className="text-default-500">12 Endpoints</small>
              <h4 className="font-bold text-large">View Metrics</h4>
            </CardHeader>
          </Card>
          <Card className="w-1/3 py-4" isPressable isHoverable>
            <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
              <small className="text-default-500">username</small>
              <h4 className="font-bold text-large">Manage Profile</h4>
            </CardHeader>
          </Card>
        </div>
      </div>
    </section>
  );
}
