# TODO: some variables needed to be set
LLVM_PATH := /home/guoy/workspace_mnt/rv/llvm-17/llvm-17/build
GNU_PATH := /home/guoy/workspace_mnt/rv/gnu-toolchain/230613/build/riscv32-gcc-ilp32f

C_FLAGS := --target=springcore32p1p0-unknown-elf -mabi=ilp32f -g -march=rv32imf -c
ASM_FLAGS := --target=springcore32p1p0-unknown-elf -g -mabi=ilp32f -march=rv32imf -c
LD_FLAGS := -g --gcc-toolchain=$(GNU_PATH) --target=riscv32-unknown-elf -T./system/link_riscv.ld -march=rv32imf -mabi=ilp32f
# set variables done

CC := $(LLVM_PATH)/bin/clang
AS := $(LLVM_PATH)/bin/clang

TARGET_EXEC ?= $(shell basename $(CURDIR)).elf

BUILD_DIR ?= ./build
SRC_DIRS := .

SRCS := $(shell find $(SRC_DIRS) -name '*.c' -or -name '*.s' -or -name '*.S')
OBJS := $(SRCS:%=$(BUILD_DIR)/%.o)

$(BUILD_DIR)/$(TARGET_EXEC): $(OBJS)
	$(CC) $(OBJS) -o $@ $(LD_FLAGS)

# assembly
$(BUILD_DIR)/%.s.o: %.s
	@$(MKDIR_P) $(dir $@)
	$(AS) $(ASM_FLAGS) $< -o $@

$(BUILD_DIR)/%.S.o: %.S
	@$(MKDIR_P) $(dir $@)
	$(AS) $(ASM_FLAGS) $< -o $@

# c source
$(BUILD_DIR)/%.c.o: %.c
	@$(MKDIR_P) $(dir $@)
	$(CC) $(C_FLAGS) $(CFLAGS) -c $< -o $@

.PHONY: clean

clean:
	$(RM) -r $(BUILD_DIR)

MKDIR_P ?= mkdir -p

