import os
import torch
import torch.nn as nn
import torch.optim as optim
from torchvision import datasets, transforms

# Dummy neural net
class SimpleNet(nn.Module):
    def __init__(self):
        super(SimpleNet, self).__init__()
        self.fc = nn.Linear(28*28, 10)

    def forward(self, x):
        return self.fc(x.view(-1, 28*28))

# Load dataset (MNIST demo)
train_loader = torch.utils.data.DataLoader(
    datasets.MNIST('/workspace/data', train=True, download=True,
                   transform=transforms.ToTensor()),
    batch_size=64, shuffle=True)

model = SimpleNet()
optimizer = optim.SGD(model.parameters(), lr=0.01)

# Train for 1 epoch (demo)
for epoch in range(1):
    for batch_idx, (data, target) in enumerate(train_loader):
        optimizer.zero_grad()
        output = model(data)
        loss = nn.CrossEntropyLoss()(output, target)
        loss.backward()
        optimizer.step()

# Save model
os.makedirs("/workspace/output", exist_ok=True)
torch.save(model.state_dict(), "/workspace/output/model.pth")

print("✅ Training complete. Model saved to /workspace/output/model.pth")
