import { useState } from "react";
import { Input } from "../ui/input";
import {
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "../ui/command";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { Checkbox } from "../ui/checkbox";
import {
  Control,
  UseFormReturn,
  useFieldArray,
  useWatch,
} from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import { FormValues } from "@/app/stock-manage/export/add-note/page";

const Total = ({ control }: { control: Control<FormValues> }) => {
  const formValues = useWatch({
    name: "ingredients",
    control,
  });
  const total = formValues.reduce(
    (acc, current) => acc + 1 * (current.quantity || 0),
    0
  );
  return <p>{total}</p>;
};

const ExportIngredient = ({
  form,
}: {
  form: UseFormReturn<FormValues, any, undefined>;
}) => {
  // const { register, handleSubmit, control, watch, getValues } = form;
  // const {
  //   fields: fieldsIngre,
  //   append: appendIngre,
  //   remove: removeIngre,
  // } = useFieldArray({
  //   control: control,
  //   name: "ingredients",
  // });
  // const [openIngre, setOpenIngre] = useState(false);
  // const [checkedIngre, setCheckedIngre] = useState(
  //   new Array(ingredients.length).fill(false)
  // );
  // const handleOnChecked = (position: number) => {
  //   const updateCheckedState = checkedIngre.map((item, index) =>
  //     index === position ? !item : item
  //   );
  //   setCheckedIngre(updateCheckedState);
  // };
  // const resetCheckedIngre = () => {
  //   setCheckedIngre(new Array(ingredients.length).fill(false));
  // };
  // const handleIngreConfirm = () => {
  //   setOpenIngre(false);
  //   checkedIngre.forEach((element, index) => {
  //     const id = ingredients.at(index)?.id!;
  //     if (element === true) {
  //       if (!fieldsIngre.find((item) => item.idIngre === id)) {
  //         appendIngre({
  //           idIngre: id,
  //           quantity: 0,
  //           price: 0,
  //         });
  //       }
  //     }
  //   });
  // };
  // return (
  //   <div className="flex flex-col">
  //     <Input
  //       className="mb-4"
  //       placeholder="Tìm nguyên liệu"
  //       onClick={() => {
  //         setOpenIngre((open) => !open);
  //         resetCheckedIngre();
  //       }}
  //     />
  //     <div className="flex pr-12 font-medium py-2 mb-2 bg-orange-100">
  //       <h2 className="basis-1/3">Tên nguyên liệu</h2>
  //       <h2 className="basis-1/3 text-center">Đơn vị</h2>
  //       <h2 className="basis-1/3 text-center">Số lượng</h2>
  //     </div>
  //     <div>
  //       {fieldsIngre.length < 1 ? (
  //         <div className="text-center py-4">Chọn sản phẩm nhập kho</div>
  //       ) : null}
  //       {fieldsIngre.map((ingre, index) => {
  //         const value = ingredients.find((item) => item.id === ingre.idIngre);
  //         return (
  //           <div key={ingre.id} className="flex items-center py-2 gap-4 ">
  //             <h2 className="basis-1/3">{value?.name}</h2>
  //             <h2 className="basis-1/3 text-center">{value?.unit.name}</h2>
  //             <Input
  //               className="basis-1/3 "
  //               type="number"
  //               min={1}
  //               max={1000}
  //               placeholder="Nhập số lượng"
  //               defaultValue={ingre.quantity}
  //               {...register(`ingredients.${index}.quantity` as const)}
  //             ></Input>
  //             <Button
  //               variant={"ghost"}
  //               className={`px-3 ml-auto`}
  //               onClick={() => {
  //                 removeIngre(index);
  //               }}
  //             >
  //               <AiOutlineClose />
  //             </Button>
  //           </div>
  //         );
  //       })}
  //     </div>
  //     <div className="flex justify-end py-2 pr-12 font-medium ">
  //       <h2 className="w-1/3">Tổng số lượng</h2>
  //       <div className="flex">
  //         <span>
  //           <Total control={control} />
  //         </span>
  //       </div>
  //     </div>
  //     <CommandDialog open={openIngre} onOpenChange={setOpenIngre}>
  //       <CommandInput placeholder="Tìm nguyên liệu" />
  //       <CommandList className="h-80">
  //         <CommandEmpty>No results found.</CommandEmpty>
  //         <CommandGroup heading="Nguyên liệu" className="h-56 overflow-y-auto">
  //           {ingredients.map((item, index) => (
  //             <CommandItem
  //               value={item.name}
  //               key={item.id}
  //               onSelect={() => {
  //                 handleOnChecked(index);
  //               }}
  //             >
  //               <div className="px-4 blur-none flex items-center gap-2 flex-1">
  //                 <Checkbox
  //                   id={item.name}
  //                   checked={checkedIngre[index]}
  //                   onCheckedChange={() => handleOnChecked(index)}
  //                 ></Checkbox>
  //                 <Label onClick={() => handleOnChecked(index)}>
  //                   {item.name}
  //                 </Label>
  //                 <Label
  //                   onClick={() => handleOnChecked(index)}
  //                   className="ml-auto"
  //                 >
  //                   {item.unit.name}
  //                 </Label>
  //               </div>
  //             </CommandItem>
  //           ))}
  //         </CommandGroup>
  //         <CommandSeparator />
  //         <CommandGroup>
  //           <div className="pt-4 pr-4 flex justify-end">
  //             <Button onClick={handleIngreConfirm}>Thêm</Button>
  //           </div>
  //         </CommandGroup>
  //       </CommandList>
  //     </CommandDialog>
  //   </div>
  // );
};

export default ExportIngredient;
