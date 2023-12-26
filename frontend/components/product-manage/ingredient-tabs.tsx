import { ingredientForChoose, measureUnits } from "@/constants";
import { UseFormReturn, useFieldArray } from "react-hook-form";
import { useState } from "react";
import {
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "../ui/command";
import { AiOutlineClose } from "react-icons/ai";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Checkbox } from "../ui/checkbox";
import { FormValues } from "@/app/product-manage/insert-product/page";

const IngredientTabs = ({
  priceId,
  form,
}: {
  priceId: string;
  form: UseFormReturn<FormValues, any, undefined>;
}) => {
  const [openIngre, setOpenIngre] = useState(false);

  const [checkedIngre, setCheckedIngre] = useState(
    new Array(ingredientForChoose.length).fill(false)
  );

  const { register, handleSubmit, control, watch, getValues } = form;
  const {
    fields: fieldsIngre,
    append: appendIngre,
    remove: removeIngre,
  } = useFieldArray({
    control: control,
    name: "ingredients",
  });
  const handleOnChecked = (position: number) => {
    const updateCheckedState = checkedIngre.map((item, index) =>
      index === position ? !item : item
    );

    setCheckedIngre(updateCheckedState);
  };

  const resetCheckedIngre = () => {
    // setCheckedIngre(new Array(ingredientForChoose.length).fill(false));
    const updated = checkedIngre.map((item, index) => {
      if (
        fieldsIngre.find(
          (ingre) =>
            ingre.idIngre === ingredientForChoose.at(index)!.id &&
            ingre.priceId === priceId
        )
      ) {
        return true;
      }
      return false;
    });
    setCheckedIngre(updated);
  };

  const handleIngreConfirm = () => {
    setOpenIngre(false);
    checkedIngre.forEach((element, index) => {
      const id = ingredientForChoose.at(index)?.id!;
      if (element === true) {
        if (
          !fieldsIngre.find(
            (item) => item.idIngre === id && item.priceId === priceId
          )
        ) {
          appendIngre({
            idIngre: id,
            amount: 0,
            priceId: priceId,
          });
        }
      } else {
        //TODO: fix the remove
        const index = fieldsIngre.findIndex(
          (item) => item.idIngre === id && item.priceId === priceId
        );
        removeIngre(index);
      }
    });
  };

  return (
    <div>
      <Input
        className="mb-4"
        placeholder="Tìm nguyên liệu"
        onClick={() => {
          setOpenIngre((open) => !open);
          resetCheckedIngre();
        }}
      />

      <div>
        {fieldsIngre.map((ingre, index) => {
          if (ingre.priceId != priceId) {
            return null;
          }
          const value = ingredientForChoose.find(
            (item) => item.id === ingre.idIngre
          );

          return (
            <div
              key={ingre.id}
              className="flex justify-between items-center mb-4"
            >
              <Label className="w-1/4">{value?.name}</Label>
              <div className="w-1/2 flex items-center gap-2">
                <Input
                  type="number"
                  min={1}
                  max={1000}
                  defaultValue={ingre.amount}
                  {...register(`ingredients.${index}.amount` as const)}
                ></Input>
                <Label className="w-8">
                  {measureUnits.find((item) => item.id === value?.unitId)?.name}
                </Label>
              </div>

              <Button
                variant={"ghost"}
                className={`self-end px-3 gap-0 `}
                onClick={() => {
                  removeIngre(index);
                }}
              >
                <AiOutlineClose />
              </Button>
            </div>
          );
        })}
      </div>
      <CommandDialog open={openIngre} onOpenChange={setOpenIngre}>
        <CommandInput placeholder="Tìm nguyên liệu" />
        <CommandList className="h-80">
          <CommandEmpty> Không tìm thấy bản ghi</CommandEmpty>
          <CommandGroup heading="Nguyên liệu" className="h-56 overflow-y-auto">
            {ingredientForChoose.map((item, index) => (
              <CommandItem
                value={item.name}
                key={item.id}
                onSelect={() => {
                  handleOnChecked(index);
                }}
              >
                <div className="px-4 blur-none flex items-center gap-2 flex-1">
                  <Checkbox
                    id={item.name}
                    checked={checkedIngre[index]}
                    onCheckedChange={() => handleOnChecked(index)}
                  ></Checkbox>
                  <Label onClick={() => handleOnChecked(index)}>
                    {item.name}
                  </Label>
                  <Label
                    onClick={() => handleOnChecked(index)}
                    className="ml-auto"
                  >
                    {measureUnits.find((unit) => unit.id === item.unitId)?.name}
                  </Label>
                </div>
              </CommandItem>
            ))}
          </CommandGroup>

          <CommandSeparator />
          <CommandGroup>
            <div className="pt-4 pr-4 flex justify-between">
              <span className="text-sm">
                {checkedIngre.filter(Boolean).length} trong 10 dòng đã được chọn
              </span>
              <Button onClick={() => handleIngreConfirm()}>Thêm</Button>
            </div>
          </CommandGroup>
        </CommandList>
      </CommandDialog>
    </div>
  );
};

export default IngredientTabs;
