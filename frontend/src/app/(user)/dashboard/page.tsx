"use client";
import {
  AddNoteIcon,
  CopyDocumentIcon,
  DeleteDocumentIcon,
  EditDocumentIcon,
} from "@/components/ui/icons";
import {
  Card,
  CardHeader,
  Listbox,
  ListboxItem,
  ListboxSection,
  cn,
} from "@nextui-org/react";

export default function Page() {
  const iconClasses =
    "text-xl text-default-500 pointer-events-none flex-shrink-0";

  return (
    <div>
      <section>
        <div className="flex items-center justify-center pt-[4rem]">
          <div className="flex flex-row max-w-lg gap-x-4 ">
            <Card className="py-4">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
                <p className="text-tiny uppercase font-bold">
                  Endpoints Protected
                </p>
                <small className="text-default-500">12 Endpoints</small>
                <h4 className="font-bold text-large">OpenAPI</h4>
              </CardHeader>
            </Card>

            <Card className="py-4">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
                <p className="text-tiny uppercase font-bold">Daily Mix</p>
                <small className="text-default-500">12 Tracks</small>
                <h4 className="font-bold text-large">Frontend Radio</h4>
              </CardHeader>
            </Card>
            <Card className="py-4">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
                <p className="text-tiny uppercase font-bold">Daily Mix</p>
                <small className="text-default-500">12 Tracks</small>
                <h4 className="font-bold text-large">Frontend Radio</h4>
              </CardHeader>
            </Card>
          </div>
        </div>
      </section>
      <section className="pl-5 pt-2">
        <ListboxWrapper>
          <Listbox variant="flat" aria-label="Listbox menu with sections">
            <ListboxSection title="Actions" showDivider>
              <ListboxItem
                key="new"
                description="Create a new file"
                startContent={<AddNoteIcon className={iconClasses} />}
              >
                New Endpoint
              </ListboxItem>
              <ListboxItem
                key="copy"
                description="Copy the file link"
                startContent={<CopyDocumentIcon className={iconClasses} />}
              >
                Copy link
              </ListboxItem>
              <ListboxItem
                key="edit"
                description="Allows you to edit the file"
                startContent={<EditDocumentIcon className={iconClasses} />}
              >
                Edit file
              </ListboxItem>
            </ListboxSection>
            <ListboxSection title="Danger zone">
              <ListboxItem
                key="delete"
                className="text-danger"
                color="danger"
                description="Permanently delete the file"
                startContent={
                  <DeleteDocumentIcon
                    className={cn(iconClasses, "text-danger")}
                  />
                }
              >
                Delete file
              </ListboxItem>
            </ListboxSection>
          </Listbox>
        </ListboxWrapper>
      </section>
    </div>
  );
}

const ListboxWrapper = ({ children }) => (
  <div className="w-full max-w-[260px] border-small px-1 py-2 rounded-small border-default-200 dark:border-default-100">
    {children}
  </div>
);
