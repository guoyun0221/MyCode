import math

def simulate(loan_amount, annual_loan_rate, initial_savings, annual_investment_return, years):
    monthly_rate = annual_loan_rate / 12
    months = years * 12

    # 等额本息：每月还款金额
    monthly_payment_equal_principal_interest = (
        loan_amount * monthly_rate * (1 + monthly_rate) ** months /
        ((1 + monthly_rate) ** months - 1)
    )

    # 等额本金：每月本金固定
    monthly_principal = loan_amount / months

    # 年度数据
    print(f"{'Year':^6} | {'方式':^6} | {'当年还款':^12} | {'累计已还':^12} | {'剩余存款':^12} | {'提前还款额':^14} | {'提前还款后存款':^16}")
    print("-" * 90)

    for method in ['等额本息', '等额本金']:
        savings = initial_savings
        total_paid = 0
        remaining_principal = loan_amount

        for year in range(1, years + 1):
            year_paid = 0

            for m in range(12):
                if method == '等额本息':
                    interest = remaining_principal * monthly_rate
                    principal = monthly_payment_equal_principal_interest - interest
                    payment = monthly_payment_equal_principal_interest
                else:
                    interest = remaining_principal * monthly_rate
                    principal = monthly_principal
                    payment = principal + interest

                remaining_principal -= principal
                year_paid += payment
                total_paid += payment
                monthly_investment_return = (1 + annual_investment_return) ** (1 / 12) - 1
                # for debug, check monthly_investment_return
                # print("monthly_investment_return:", monthly_investment_return)
                savings *= (1 + monthly_investment_return)
                savings -= payment

            # 提前还款模拟
            early_repay_amount = remaining_principal  # 不计手续费
            savings_after_early_repay = savings - early_repay_amount

            print(f"{year:^6} | {method:^6} | {year_paid:>12,.2f} | {total_paid:>12,.2f} | {savings:>12,.2f} | {early_repay_amount:>14,.2f} | {savings_after_early_repay:>16,.2f}")

        print("-" * 90)

# 示例调用
simulate(
    loan_amount=850000,
    annual_loan_rate=0.028,
    initial_savings=850000,
    annual_investment_return=0.07,
    years=30
)
