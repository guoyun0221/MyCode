import matplotlib.pyplot as plt

def simulate_retirement(A, R, I, T, S):
    years = list(range(T + 1))
    assets = [S]
    for n in years[:-1]:
        inflation_adjusted_spending = A * (1 + I) ** n
        year_end_assets = assets[-1] - inflation_adjusted_spending + assets[-1] * R
        assets.append(max(year_end_assets, 0))  # Prevent assets from becoming negative
        print("year:", n, ", spend:", inflation_adjusted_spending, ", year_end_assets:", year_end_assets)

    return years, assets

# Parameter settings
A = 6  # Annual expenditure (in ten thousand yuan) # 当前年支出 (万元)
R = 0.07  # Annual investment return rate # 年投资回报率
I = 0.04  # Annual inflation rate # 年通胀率 
T = 30  # Number of retirement years # 预计退休生活年限
S = 130  # Initial savings (in ten thousand yuan) # 初始资产 (万元)

# Simulation
years, assets = simulate_retirement(A, R, I, T, S)

# Visualization
plt.plot(years, assets, label="Asset Change")
plt.axhline(0, color='red', linestyle='--', label="Asset Depletion") # 资产耗尽
plt.xlabel("Retirement Years")
plt.ylabel("Assets (in ten thousand yuan)")
plt.title("Retirement Savings Simulation")
plt.legend()
plt.show()
